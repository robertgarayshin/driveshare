package main

import (
	"github.com/jinzhu/gorm"
	"github.com/robertgarayshin/driveshare/pkg/model/user"
	"log"
	"net"
	"net/http"
	"user/config"
	"user/domain"
	"user/repository"
	grpc2 "user/transport/grpc"
	"user/transport/rest"
	"user/usecase"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	grpcServer := grpc.NewServer()
	userHandler := grpc2.NewUserHandler(userUsecase)
	user.RegisterUserServiceServer(grpcServer, userHandler)
	reflection.Register(grpcServer)

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("gRPC server listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	router := mux.NewRouter()
	restHandler := rest.NewUserHandler(userUsecase)
	router.HandleFunc("/users", restHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", restHandler.GetUser).Methods("GET")

	log.Println("HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

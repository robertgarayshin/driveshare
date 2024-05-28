package main

import (
	"gateway/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	userConn, err := grpc.Dial("user-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer userConn.Close()

	carConn, err := grpc.Dial("car-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer carConn.Close()

	bookingConn, err := grpc.Dial("booking-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer bookingConn.Close()

	gatewayHandler := rest.NewGatewayHandler(userConn, carConn, bookingConn)

	router := mux.NewRouter()
	router.HandleFunc("/users", gatewayHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", gatewayHandler.GetUser).Methods("GET")
	// Add routes for Car and Booking

	log.Println("Gateway server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"fmt"
	"github.com/robertgarayshin/driveshare/proto/users"
	"google.golang.org/grpc"
	"log"
	"net"
	"rents/presenter"
)

const port = 50522

func main() {
	log.Println("Starting rents microservice...")
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	users.RegisterUsersServer(grpcServer, new(presenter.RentsServer)) //todo после подтягивания зависимостей добавить
	log.Printf("Service started on port :%d", port)
	grpcServer.Serve(listener)

}

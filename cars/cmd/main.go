package main

import (
	"cars/presenter"
	"fmt"
	"github.com/robertgarayshin/driveshare/proto/cars"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 50523

func main() {
	log.Println("Starting rents microservice...")

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	cars.RegisterCarsServer(grpcServer, new(presenter.CarServer)) //todo после подтягивания зависимостей добавить
	log.Printf("Service started on port :%d", port)

	grpcServer.Serve(listener)
}

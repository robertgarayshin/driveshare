package main

import (
	"fmt"
	"github.com/robertgarayshin/driveshare/proto"
	"github.com/robertgarayshin/driveshare/users/presenter"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 50521

func main() {
	log.Println("Starting users microservice...")
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterUserServer(grpcServer, new(presenter.UserServer))
	log.Printf("Service started on port :%d", port)
	grpcServer.Serve(listener)
}

package main

import (
	"github.com/robertgarayshin/driveshare/http"
	"github.com/robertgarayshin/driveshare/http/api"
	pb "github.com/robertgarayshin/driveshare/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	serverAddr = "localhost:50521"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)

	server := new(http.Server)
	handler := api.NewHandler(client)
	server.Run(handler.InitRoutes(), "8080")
}

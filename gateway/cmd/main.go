package main

import (
	"gateway/rest"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	userConn, err := grpc.NewClient("user-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer userConn.Close()

	carConn, err := grpc.NewClient("car-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer carConn.Close()

	bookingConn, err := grpc.NewClient("booking-service:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer bookingConn.Close()

	gatewayHandler := rest.NewGatewayHandler(userConn, carConn, bookingConn)

	router := gatewayHandler.RegisterRoutes()

	log.Println("Gateway server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

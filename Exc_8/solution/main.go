package main

import (
	"exc8/client"
	"exc8/server"
	"log"
	"time"
)

func main() {
	go func() {
		// todo start server
		if err := server.StartGrpcServer(); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
	time.Sleep(1 * time.Second)
	// todo start client
	grpcClient, err := client.NewGrpcClient()
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer grpcClient.Close()

	if err := grpcClient.Run(); err != nil {
		log.Fatalf("Client run failed: %v", err)
	}

	println("Orders complete!")
}

package main

import (
	"finance/internal/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	server := server.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server.Port()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// Register gRPC services here

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

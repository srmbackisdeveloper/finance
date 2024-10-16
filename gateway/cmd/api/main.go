package main

import (
	"gateway/internal/server"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	grpcAddress := os.Getenv("USERS_GRPC_ADDRESS")
	log.Printf("Connecting to USERS service at: %s", grpcAddress)
	
	srv, err := server.NewServer(grpcAddress)
    if err != nil {
        log.Fatalf("GATEWAY service: Failed to create server: %s", err)
    }

	err = srv.ListenAndServe()
    if err != nil {
        log.Fatalf("GATEWAY service: Failed to start server: %s", err)
    }
}

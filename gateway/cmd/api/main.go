package main

import (
	"gateway/internal/server"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("GATEWAY service: Failed to create server: %s", err)
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("GATEWAY service: Failed to start server: %s", err)
	}
}

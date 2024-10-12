package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	proto "gateway/internal/proto"

	"google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
    port       int
    grpcClient proto.UserServiceClient
}

func NewServer(grpcAddress string) (*http.Server, error) {
    port, _ := strconv.Atoi(os.Getenv("PORT"))

    conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
    if err != nil {
        return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
    }

    grpcClient := proto.NewUserServiceClient(conn)

    newServer := &Server{
        port:       port,
        grpcClient: grpcClient,
    }

    httpServer := &http.Server{
        Addr:         fmt.Sprintf(":%d", newServer.port),
        Handler:      newServer.RegisterRoutes(),
        IdleTimeout:  time.Minute,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

    return httpServer, nil
}

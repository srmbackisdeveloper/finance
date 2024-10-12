package server

import (
	"fmt"
	"net"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"

	"users/internal/database"
	pb "users/internal/proto"

	"google.golang.org/grpc"
)

type Server struct {
	port int
	pb.UnimplementedUserServiceServer
	db   database.Service
}

func NewServer() *grpc.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, NewServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", NewServer.port))
	if err != nil {
		panic(fmt.Sprintf("failed to listen on port %d: %v", NewServer.port, err))
	}

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			panic(fmt.Sprintf("failed to serve gRPC server: %v", err))
		}
	}()

	return grpcServer
}

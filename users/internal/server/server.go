package server

import (
	"fmt"
	"github.com/IBM/sarama"
	_ "github.com/joho/godotenv/autoload"
	"net"
	"os"

	"users/internal/database"
	pb "users/internal/proto"

	"google.golang.org/grpc"
)

type Server struct {
	port string
	pb.UnimplementedUserServiceServer
	db       database.Service
	producer sarama.SyncProducer
}

func NewServer() *grpc.Server {
	port := os.Getenv("PORT")

	producer, err := NewKafkaProducer()
	if err != nil {
		panic(fmt.Sprintf("failed to create Kafka producer: %v", err))
	}

	NewServer := &Server{
		port:     port,
		db:       database.New(),
		producer: producer,
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)
	pb.RegisterUserServiceServer(grpcServer, NewServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", NewServer.port))
	if err != nil {
		panic(fmt.Sprintf("failed to listen on port %s: %v", NewServer.port, err))
	}

	fmt.Printf("gRPC server USERS is starting on port %s...\n", NewServer.port)

	if err = grpcServer.Serve(listener); err != nil {
		panic(fmt.Sprintf("failed to serve gRPC server: %v", err))
	}

	return grpcServer
}

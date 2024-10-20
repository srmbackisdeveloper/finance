package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	proto "gateway/internal/proto"

	"google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port            int
	usersGrpcClient proto.UserServiceClient
	// TODO: payment service connection
	// paymentGrpcClient proto.PaymentServiceClient

	// TODO: finance service connection
	// financeGrpcClient proto.FinanceServiceClient
}

func NewServer() (*http.Server, error) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	grpcAddress := os.Getenv("USERS_GRPC_ADDRESS")
	log.Printf("Connecting to USERS service at: %s", grpcAddress)

	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	log.Println("successful connection with USERS service")

	usersGrpcClient := proto.NewUserServiceClient(conn)

	newServer := &Server{
		port:            port,
		usersGrpcClient: usersGrpcClient,
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

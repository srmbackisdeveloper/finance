package server

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("Received gRPC method: %s, with request: %v", info.FullMethod, req)

	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("Error in method: %s, error: %v", info.FullMethod, err)
	} else {
		log.Printf("Successfully handled method: %s", info.FullMethod)
	}

	return resp, err
}

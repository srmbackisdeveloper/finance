package server

import (
	"context"
	pb "users/internal/proto"
)


func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		Id:   1,
		Name: "name_1",
	}, nil
}


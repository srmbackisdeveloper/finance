package server

import (
	"context"
	pb "users/internal/proto"
)



func (s *Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) { 
	return &pb.RegisterUserResponse{Message: "msg"}, nil
}

func (s *Server) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerifyUserResponse, error) { 
	return &pb.VerifyUserResponse{Message: "msg"}, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) { 
	return &pb.LoginUserResponse{
		AccessToken: "access",
		RefreshToken: "refresh",
	}, nil
}

func (s *Server) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) { 
	return &pb.RefreshTokenResponse{
		AccessToken: "access",
		RefreshToken: "refresh",
	}, nil
}

func (s *Server) Health(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) { 
	return &pb.HealthResponse{Message: "good"}, nil
}
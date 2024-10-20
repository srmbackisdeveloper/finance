package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
	"users/internal/helpers"
	"users/internal/models"
	pb "users/internal/proto"
)

func (s *Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	// Log the incoming registration request
	log.Printf("Registering user with email: %s", req.Email)

	// Check if user already exists
	user, err := s.db.GetUserByEmail(req.Email)
	if err != nil {
		log.Printf("Error checking for existing user: %v", err)
		return nil, status.Error(codes.Internal, "could not check for existing")
	}
	if user != nil {
		log.Printf("User with email %s already exists", req.Email)
		return nil, status.Error(codes.AlreadyExists, "account with this email already exists")
	}

	// Hash password
	log.Printf("Hashing password for user: %s", req.Email)
	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		log.Printf("Error hashing password for user %s: %v", req.Email, err)
		return nil, status.Error(codes.Internal, "could not encrypt the password")
	}

	// Generate verification code
	log.Printf("Generating verification code for user: %s", req.Email)
	verificationCode, err := helpers.GenerateVerificationCode()
	if err != nil {
		log.Printf("Error generating verification code for user %s: %v", req.Email, err)
		return nil, status.Error(codes.Internal, "could not generate a verification code")
	}

	// Create new user
	log.Printf("Creating new user: %s", req.Email)
	if err = s.db.CreateUser(&models.User{
		Name:                   "Default Name",
		Email:                  req.Email,
		Password:               hashPassword,
		IsPremium:              false,
		IsVerified:             false,
		VerificationCode:       verificationCode,
		VerificationValidUntil: time.Now().Add(15 * time.Minute),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}); err != nil {
		log.Printf("Error creating user %s: %v", req.Email, err)
		return nil, status.Error(codes.Internal, "could not create a user")
	}

	// Send verification email
	log.Printf("Sending verification email to: %s", req.Email)
	if err = s.SendUserRegistered(req.Email, verificationCode); err != nil {
		log.Printf("Failed to send verification message to %s: %v", req.Email, err)
		return nil, status.Error(codes.Internal, "could not send verification message")
	}

	log.Printf("Successfully registered user: %s", req.Email)
	return &pb.RegisterUserResponse{Message: "Registration successful"}, nil
}

func (s *Server) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerifyUserResponse, error) {
	user, err := s.db.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		return nil, status.Error(codes.NotFound, "could not find a user")
	}

	if user.IsVerified {
		return nil, status.Error(codes.AlreadyExists, "already verified")
	}

	if user.VerificationCode != req.Code {
		return nil, status.Error(codes.InvalidArgument, "incorrect verification code")
	}

	if time.Now().After(user.VerificationValidUntil) {
		verificationCode, err := helpers.GenerateVerificationCode()
		if err != nil {
			return nil, status.Error(codes.Internal, "could not generate a verification code")
		}

		if err = s.SendUserRegistered(req.Email, verificationCode); err != nil {
			log.Printf("Failed to send verification message: %v", err)
			return nil, status.Error(codes.Internal, "could not send verification message")
		}

		return nil, status.Error(codes.Unknown, "verification code expired, new one will be sent to email")
	}

	err = s.db.VerifyUser(user.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "could not verify user")
	}

	return &pb.VerifyUserResponse{Message: "success"}, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := s.db.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		return nil, status.Error(codes.NotFound, "could not find a user")
	}

	if !user.IsVerified {
		return nil, status.Error(codes.NotFound, "user is not verified")
	}

	if err := helpers.CheckPassword(req.Password, user.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid login or password")
	}

	accessToken, err := helpers.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate access token")
	}

	refreshToken, err := helpers.GenerateRefreshToken()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate refresh token")
	}

	err = s.db.StoreRefreshToken(&models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to store refresh token")
	}

	return &pb.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	storedToken, err := s.db.GetRefreshToken(req.RefreshToken)
	if err != nil || storedToken == nil {
		return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
	}

	if time.Now().After(storedToken.ExpiresAt) {
		return nil, status.Error(codes.Unauthenticated, "refresh token expired")
	}

	newAccessToken, err := helpers.GenerateAccessToken(storedToken.UserID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate access token")
	}

	newRefreshToken, err := helpers.GenerateRefreshToken()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate refresh token")
	}

	err = s.db.UpdateRefreshToken(storedToken.Token, newRefreshToken)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to store new refresh token")
	}

	return &pb.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

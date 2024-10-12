package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "gateway/internal/proto"
)

func (s *Server) registerHandler(ctx *gin.Context) {
    var regPayload struct {
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&regPayload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    grpcReq := &pb.RegisterUserRequest{
        Email:    regPayload.Email,
        Password: regPayload.Password,
    }

    grpcResp, err := s.grpcClient.RegisterUser(context.Background(), grpcReq)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": grpcResp.Message,
    })
}

func (s *Server) verifyHandler(ctx *gin.Context) {
    var verifyPayload struct {
        Email string `json:"email" binding:"required"`
        Code  string `json:"code" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&verifyPayload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    grpcReq := &pb.VerifyUserRequest{
        Email: verifyPayload.Email,
        Code:  verifyPayload.Code,
    }

    grpcResp, err := s.grpcClient.VerifyUser(context.Background(), grpcReq)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": grpcResp.Message,
    })
}

func (s *Server) loginHandler(ctx *gin.Context) {
    var loginPayload struct {
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&loginPayload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    grpcReq := &pb.LoginUserRequest{
        Email:    loginPayload.Email,
        Password: loginPayload.Password,
    }

    grpcResp, err := s.grpcClient.LoginUser(context.Background(), grpcReq)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "access_token":  grpcResp.AccessToken,
        "refresh_token": grpcResp.RefreshToken,
    })
}

func (s *Server) refreshTokenHandler(ctx *gin.Context) {
    var tokenPayload struct {
        RefreshToken string `json:"refresh_token" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&tokenPayload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    grpcReq := &pb.RefreshTokenRequest{
        RefreshToken: tokenPayload.RefreshToken,
    }

    grpcResp, err := s.grpcClient.RefreshToken(context.Background(), grpcReq)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "access_token":  grpcResp.AccessToken,
        "refresh_token": grpcResp.RefreshToken,
    })
}

func (s *Server) getUserHandler(ctx *gin.Context) {
    userID, exists := ctx.Get("userID")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    grpcReq := &pb.GetUserRequest{
        UserId: int32(userID.(int)), // grpc expects int32
    }

    grpcResp, err := s.grpcClient.GetUser(context.Background(), grpcReq)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "id":         grpcResp.Id,
        "name":       grpcResp.Name,
        "email":      grpcResp.Email,
        "created_at": grpcResp.CreatedAt,
        "updated_at": grpcResp.UpdatedAt,
    })
}

func (s *Server) updateUserHandler(ctx *gin.Context) { 
	ctx.JSON(http.StatusOK, gin.H{
        "warning": "Under development",
    })
}

func (s *Server) deleteUserHandler(ctx *gin.Context) { 
	ctx.JSON(http.StatusOK, gin.H{
        "warning": "Under development",
    })
}
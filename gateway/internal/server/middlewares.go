package server

import (
	"gateway/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Server) authMiddleware(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		ctx.Abort()
		return
	}

	// parse and validate the JWT token
	token, err := helpers.ValidateAccessToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		ctx.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// sub = userID
        userID := int(claims["sub"].(float64)) 
        ctx.Set("userID", userID)
    }

	ctx.Next()
}
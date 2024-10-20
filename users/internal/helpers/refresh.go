package helpers

import (
	"encoding/base64"
	"fmt"
	"math/rand"
)

func GenerateRefreshToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random refresh token: %w", err)
	}

	refreshToken := base64.URLEncoding.EncodeToString(tokenBytes)
	return refreshToken, nil
}

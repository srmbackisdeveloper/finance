package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateVerificationCode() (string, error) {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000)), nil
}

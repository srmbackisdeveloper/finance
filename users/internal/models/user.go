package models

import "time"

type User struct {
	ID                 int64     `json:"id"`
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	IsVerified         bool      `json:"is_verified"`
	VerificationCode   string    `json:"verification_code"`
	VerificationSentAt time.Time `json:"verification_sent_at"`
	IsPremium 		   bool      `json:"is_premium"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

package models

import "time"

type User struct {
	ID                     int64     `json:"id"`
	Name                   string    `json:"name"`
	Email                  string    `json:"email"`
	Password               string    `json:"password"`
	IsPremium              bool      `json:"is_premium"`
	IsVerified             bool      `json:"is_verified"`
	VerificationCode       string    `json:"verification_code"`
	VerificationValidUntil time.Time `json:"verification_valid_until"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

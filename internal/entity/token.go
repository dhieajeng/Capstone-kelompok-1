package entity

import (
	"time"

	"github.com/google/uuid"
)

type TokenAction string

const (
	Register       TokenAction = "register"
	ForgotPassword TokenAction = "forgot-password"
)

type Token struct {
	ID        uuid.UUID   `json:"id"`
	Email     string      `json:"email"`
	Action    TokenAction `json:"action"`
	ExpiresAt time.Time   `json:"expires_at" sql:"expires_at"`
	Auditable
}

func NewToken(email string, action TokenAction) *Token {
	return &Token{
		ID:        uuid.New(),
		Email:     email,
		Action:    action,
		ExpiresAt: time.Now().Add(time.Minute * 15),
		Auditable: NewAuditable(),
	}
}

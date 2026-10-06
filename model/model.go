package model

import (
	"time"
)

type Claims struct {
	ID       int
	Username string
}
type contextKey string

const ClaimsKey contextKey = "claims"

type Article struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Judul     string    `json:"judul" binding:"required,min=3,max=255"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

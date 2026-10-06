package model

import (
	"time"
)

type Registrasi struct {
	Username        string `json:"username" validate:"required,min=3,max=15"`
	Email           string `json:"email" validate:"required,email"`
	NoTelepon       string `json:"no_telepon" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type TEst interface{
	tes()
}

type Login struct {
	Login    string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username" validate:"required,min=3,max=15"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required,min=8"`
	NoTelepon string `json:"no_telepon" validate:"required"`
}
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

package model

type Registrasi struct {
	Username        string `json:"username" validate:"required,min=3,max=15"`
	Email           string `json:"email" validate:"required,email"`
	NoTelepon       string `json:"no_telepon" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type Login struct {
	ID       int    `json:"id"`
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username" validate:"required,min=3,max=15"`
	Email     string `json:"email" validate:"required,email"`
	NoTelepon string `json:"no_telepon" validate:"required"`
}

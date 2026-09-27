package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Gagal membaca .env")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

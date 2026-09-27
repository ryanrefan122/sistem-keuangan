package main

import (
	"main/config"
	"main/config/database"
	"main/router"
)

func main() {
	config.LoadEnv()
	database.Connect()
	router.Server()
}

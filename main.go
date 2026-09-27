package main

import (
	"main/config/database"
	"main/router"
)

func main() {
	database.Connect()
	router.Server()
}

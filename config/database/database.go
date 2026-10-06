package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	dsn := user + ":" + password +
		"@tcp(" + host + ":" + port + ")/login?parseTime=true" + database
	db, err := sql.Open(
		"mysql",
		dsn,
	)
	if err != nil {
		fmt.Println("sql tak terhubung")
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("sql bermasalah")
		return
	}
	DB = db
	fmt.Println("terhubung")
}

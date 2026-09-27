package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	db, err := sql.Open(
		"mysql",
		"root:kinan123@tcp(localhost:3306)/login?parseTime=true",
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

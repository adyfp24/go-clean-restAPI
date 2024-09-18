package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	dsn := "root:@tcp(127.0.0.1:3306)/db_clean_go"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Errorf("failed to connect database : %v", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Errorf("failed to ping database : %v", err)
	}

	DB = db
	log.Println("database connection succesfull")
	return nil
}

package database

import (
	"database/sql"
	"fmt"
	"go-clean-restAPI/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	dbConfig := config.AppConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
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

package migrations

import (
	"fmt"
	"go-clean-restAPI/internal/entity"
	"go-clean-restAPI/pkg/database"
)

func RunMigration(){

	if err := database.InitDB(); err != nil {
		panic(err)
	}

	err := database.DB.AutoMigrate(
		&entity.User{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("migration success")
}
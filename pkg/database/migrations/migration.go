package migrations

import (
	"fmt"
	"go-clean-restAPI/internal/entity"
	"go-clean-restAPI/pkg/database"
	"go-clean-restAPI/pkg/database/seeders"
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

	seeders.AddUserSeeder(database.DB)	

	fmt.Println("migration success")
}
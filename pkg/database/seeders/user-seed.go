package seeders

import (
	"fmt"
	"go-clean-restAPI/internal/entity"
	"log"

	"gorm.io/gorm"
)

func AddUserSeeder(db *gorm.DB) {
	users := []entity.User{
		{Name: "John Doe", Age: 30, Address: "123 Main St", Role: "Admin"},
		{Name: "Jane Smith", Age: 25, Address: "456 Oak St", Role: "User"},
		{Name: "Alex Johnson", Age: 28, Address: "789 Pine St", Role: "User"},
	}

	if result := db.Where("name = ?", users[0].Name).First(&entity.User{}); result.Error != nil {
		if err := db.Create(&users).Error; err != nil {
			log.Fatalf("Failed to seed users: %v", err)
		} else {
			fmt.Println("Users seeded successfully.")
		}
	} else {
		fmt.Println("Users already seeded.")
	}
}
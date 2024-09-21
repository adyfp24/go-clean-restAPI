package main

import (
	"fmt"
	"go-clean-restAPI/pkg/database"
	"log"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	fmt.Println("wkwkw golang ni bos")
}

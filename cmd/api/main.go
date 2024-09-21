package main

import (
	"fmt"
	"go-clean-restAPI/internal/handler"
	"go-clean-restAPI/internal/repository"
	"go-clean-restAPI/internal/route"
	"go-clean-restAPI/internal/usecase"
	"go-clean-restAPI/pkg/database"
	"log"
	"net/http"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer database.DB.Close()

	repositories := repository.NewRepositories(database.DB)
	usecases := usecase.NewUseCases(repositories)
	handlers := handler.NewHandlers(usecases)

	r := route.InitRoute(handlers)

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("wkwkw golang ni bos")
}

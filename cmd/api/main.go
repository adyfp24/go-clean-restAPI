package main

import (
	"go-clean-restAPI/config"
	"go-clean-restAPI/internal/handler"
	"go-clean-restAPI/internal/repository"
	"go-clean-restAPI/internal/route"
	"go-clean-restAPI/internal/usecase"
	"go-clean-restAPI/pkg/database"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	if err := database.InitDB(); err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	repositories := repository.NewRepositories(database.DB)
	usecases := usecase.NewUseCases(repositories)
	handlers := handler.NewHandlers(usecases)

	r := route.InitRoute(handlers)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

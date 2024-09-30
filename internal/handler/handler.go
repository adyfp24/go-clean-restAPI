package handler

import "go-clean-restAPI/internal/usecase"

type Handlers struct {
	UserHandler UserHandler
	AuthHandler AuthHandler
}

func NewHandlers(u *usecase.UseCases) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(u.UserUsecase),
	}
}

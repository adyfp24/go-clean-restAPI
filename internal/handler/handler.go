package handler

import "go-clean-restAPI/internal/usecase"

type Handlers struct {
	UserHandler UserHandler
}

func NewHandlers(u *usecase.UseCases) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(u.UserUsecase),
	}
}

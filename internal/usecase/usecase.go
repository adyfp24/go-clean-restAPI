package usecase

import "go-clean-restAPI/internal/repository"

type UseCases struct {
	UserUsecase UserUsecase
	AuthUsecase AuthUsecase
}

func NewUseCases(r *repository.Repositories) *UseCases {
	return &UseCases{
		UserUsecase: NewUserUsecase(r.UserRepo),
		AuthUsecase: NewAuthUsecase(r.AuthRepo),
	}
}

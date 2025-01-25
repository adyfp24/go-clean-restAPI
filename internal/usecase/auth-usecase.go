package usecase

import (
	"go-clean-restAPI/internal/repository"
)

type AuthUsecase interface {
	Login()
	Register()
}

type authUsecase struct {
	authRepo repository.AuthRepo
}

func NewAuthUsecase(authRepo repository.AuthRepo) AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (a *authUsecase) Login() {}

func (a *authUsecase) Register() {}

package usecase

import (
	"go-clean-restAPI/internal/entity"
	"go-clean-restAPI/internal/repository"
)

type UserUsecase interface {
	GetAllUser() ([]entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) GetAllUser() ([]entity.User, error) {
	return u.userRepo.GetAllUser()
}

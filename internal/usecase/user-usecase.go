package usecase

import (
	"go-clean-restAPI/internal/dto"
	"go-clean-restAPI/internal/repository"
)

type UserUsecase interface {
	GetAllUser() ([]dto.UserDTO, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) GetAllUser() ([]dto.UserDTO, error) {
	users, err := u.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	usersDTO := make([]dto.UserDTO, len(users))
	return usersDTO, nil
}

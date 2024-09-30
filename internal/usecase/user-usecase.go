package usecase

import (
	"go-clean-restAPI/internal/dto"
	"go-clean-restAPI/internal/repository"
	"log"
)

type UserUsecase interface {
	GetAllUser() ([]dto.UserDTO, error)
	CreateUser(user dto.UserDTO) (dto.UserDTO, error)
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
	users, err := u.userRepo.GetAll()

	if err != nil {
		return nil, err
	}

	usersDTO := make([]dto.UserDTO, len(users))
	for i, user := range users {
		usersDTO[i] = dto.UserDTO{
			ID:      user.ID,
			Name:    user.Name,
			Age:     user.Age,
			Address: user.Address,
		}
	}

	return usersDTO, nil
}

func (u *userUsecase) CreateUser(user dto.UserDTO) (dto.UserDTO, error) {
	newUser, err := u.userRepo.Create(user)
	if err != nil {
		log.Printf("error while create new user: %v", err.Error())
		return dto.UserDTO{}, err
	}

	newUserDTO := dto.UserDTO{
		ID:      newUser.ID,
		Name:    newUser.Name,
		Age:     newUser.Age,
		Address: newUser.Address,
		Role:    newUser.Role,
	}

	return newUserDTO, nil
}

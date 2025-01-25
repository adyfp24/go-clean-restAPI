package repository

import (
	"go-clean-restAPI/internal/dto"
	"go-clean-restAPI/internal/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetAll() ([]entity.User, error)
	Create(newUser dto.UserRequest) (entity.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetAll() ([]entity.User, error) {
    var users []entity.User
    result := r.db.Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }
    return users, nil
}

func (r *userRepo) Create(newUser dto.UserRequest) (entity.User, error) {
    user := entity.User{
        Name:    newUser.Name,
        Age:     newUser.Age,
        Address: newUser.Address,
        Role:    newUser.Role,
    }
    result := r.db.Create(&user)
    if result.Error != nil {
        return entity.User{}, result.Error
    }
    return user, nil
}


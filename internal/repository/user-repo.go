package repository

import (
	"database/sql"
	"go-clean-restAPI/internal/entity"
)

type UserRepo interface {
	GetAllUser() ([]entity.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetAllUser() ([]entity.User, error) {
	return []entity.User{}, nil
}

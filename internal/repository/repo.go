package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo UserRepo
	AuthRepo AuthRepo
}

func NewRepositories(DB *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo: NewUserRepo(DB),
		AuthRepo: NewAuthRepo(DB),
	}
}

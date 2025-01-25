package repository

import (
	"gorm.io/gorm"
)

type AuthRepo interface {
	Login()
	Register()
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) Login(){
	
}
func (r *authRepo) Register(){

}
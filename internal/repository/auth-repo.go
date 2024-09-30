package repository

import (
	"database/sql"
)

type AuthRepo interface {
	Login()
	Register()
}

type authRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) Login() {

}

func (r *authRepo) Register() {}

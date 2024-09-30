package repository

import "database/sql"

type Repositories struct {
	UserRepo UserRepo
	AuthRepo AuthRepo
}

func NewRepositories(DB *sql.DB) *Repositories {
	return &Repositories{
		UserRepo: NewUserRepo(DB),
		AuthRepo: NewAuthRepo(DB),
	}
}

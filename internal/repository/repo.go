package repository

import "database/sql"

type Repositories struct {
	UserRepo UserRepo
}

func NewRepositories(DB *sql.DB) *Repositories {
	return &Repositories{
		UserRepo: NewUserRepo(DB),
	}
}

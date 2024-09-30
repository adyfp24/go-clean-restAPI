package repository

import (
	"database/sql"
	"go-clean-restAPI/internal/entity"
	"log"
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

	var query = "SELECT id, name, age, address, role FROM users"

	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error where execute query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Role)
		if err != nil {
			log.Printf("Error scanning user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
	}

	return users, nil
}

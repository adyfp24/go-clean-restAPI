package repository

import (
	"database/sql"
	"go-clean-restAPI/internal/dto"
	"go-clean-restAPI/internal/entity"
	"log"
)

type UserRepo interface {
	GetAll() ([]entity.User, error)
	Create(newUser dto.UserDTO) (entity.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetAll() ([]entity.User, error) {

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

func (r *userRepo) Create(newUser dto.UserDTO) (entity.User, error) {
	var query = "INSERT INTO users (name, age, address, role) VALUES (?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Printf("Error while preparing statement: %v", err)
		return entity.User{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(newUser.Name, newUser.Age, newUser.Address, newUser.Role)
	if err != nil {
		log.Printf("Error while executing statement: %v", err)
		return entity.User{}, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error while executing statement: %v", err)
		return entity.User{}, err
	}

	createdUser := entity.User{
		ID:      int(lastInsertId),
		Name:    newUser.Name,
		Age:     newUser.Age,
		Address: newUser.Address,
		Role:    newUser.Role,
	}

	return createdUser, nil

}

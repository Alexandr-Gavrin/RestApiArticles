package storage

import (
	"ServerAndDB2/internal/app/models"
	"fmt"
	"log"
)

// Instance of User repository (model interface)
type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

// Create User in DB
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES($1, $2) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

// Find user by login

func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	founded := false
	if err != nil {
		return nil, founded, err
	}
	var userFinded *models.User
	for _, v := range users {
		if v.Login == login {
			founded = true
			userFinded = v
			break
		}
	}
	return userFinded, founded, nil
}

// Select all users in DB
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}

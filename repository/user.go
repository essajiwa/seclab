package repository

import (
	"fmt"
	"log"
	"seclab/model"
)

func (r *Repository) FindUserByID(id int) (model.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)
	log.Println("Query: ", query)

	data := r.db.QueryRow(query)
	var user model.User
	err := data.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) Login(username, password string) (model.User, error) {
	query := "SELECT * FROM users WHERE username = '" + username + "' AND password = '" + password + "'"
	log.Println("Query: ", query)

	data := r.db.QueryRow(query)
	var user model.User
	err := data.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}

	return user, nil
}

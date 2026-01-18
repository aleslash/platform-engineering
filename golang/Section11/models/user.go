package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	hashedPassword := utils.HashPassword(u.Password)

	result, err := db.DB.Exec(query, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	resultID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = resultID

	return err
}

func (u *User) ValidadeCredentials() error {
	query := "SELECT password, id FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var hashedPassword string
	var userId int64
	err := row.Scan(&hashedPassword, &userId)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, hashedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	u.ID = userId
	return nil
}

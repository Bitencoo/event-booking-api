package models

import (
	"github.com/Bitencoo/event-booking-api.git/db"
	"github.com/Bitencoo/event-booking-api.git/utils"
)

type User struct {
	ID int64
	Email string `binding: required`
	Password string `binding: required`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUEs (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hPassword, err := utils.HashPassword((*u).Password)
	
	if err != nil {
		return err
	}

	(*u).Password = hPassword

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	(*u).ID = id

	return err
}

func (u *User) ValidateCredentials() error {
	user, err := findUserByEmail(u.Email)

	if err != nil {
		return err
	}

	(*u).ID = user.ID
	err = utils.ValidadePassword(u.Password, user.Password)

	return err
}

func findUserByEmail(email string) (*User, error) {
	query := `SELECT * FROM USERS WHERE email = ?`

	row := db.DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, err
}
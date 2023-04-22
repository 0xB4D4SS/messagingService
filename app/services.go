package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type AuthService interface {
	Register(string, string) (string, error)
	Login(string, string) (string, error)
	Logout(string) (string, error)
}

type authService struct{}

func (authService) Register(login string, password string) (string, error) {
	if login == "" || password == "" {
		return "", ErrEmpty
	}

	db, conErr := sql.Open("mysql", dbConfig)

	if conErr != nil {
		return "", conErr
	}

	hash := GenerateSHA256Hash(password)
	token := GenerateSecureToken(tokenDefaultLength)
	hash = fmt.Sprintf("%x", hash)
	_, err := db.Query(
		"insert into users (login, password, token) values (?, ?, ?)",
		login,
		hash,
		token,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (authService) Login(login string, password string) (string, error) {
	if login == "" || password == "" {
		return "", ErrEmpty
	}

	db, conErr := sql.Open("mysql", dbConfig)

	if conErr != nil {
		return "", conErr
	}

	hash := GenerateSHA256Hash(password)
	hash = fmt.Sprintf("%x", hash)
	rows, err := db.Query(
		"select * from users where `login` = ? and `password` = ?",
		login,
		hash,
	)

	if err != nil {
		return "", err
	}

	if rows.Next() != false {
		u := User{}
		scanErr := rows.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

		if scanErr != nil {
			return "", scanErr
		}

		if u.Token != nil {
			return *u.Token, nil
		}

		token := GenerateSecureToken(tokenDefaultLength)
		_, err := db.Query(
			"update users set token = ? where `login` = ? and `password` = ?",
			token,
			login,
			string(hash[:]),
		)

		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", ErrNotFound
}

func (authService) Logout(token string) (string, error) {
	if token == "" {
		return "", ErrEmpty
	}

	db, conErr := sql.Open("mysql", dbConfig)

	if conErr != nil {
		return "", conErr
	}

	rows, err := db.Query(
		"select * from users where `token` = ?",
		token,
	)

	if err != nil {
		return "", err
	}

	if rows.Next() != false {
		u := User{}
		scanErr := rows.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

		if scanErr != nil {
			return "", scanErr
		}

		_, err := db.Query("update users set token = null where `token` = ?", token)

		if err != nil {
			return "", err
		}

		return "Logged out", nil
	}

	return "", ErrNotFound
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

// ErrNotFound is returned when no data found in db
var ErrNotFound = errors.New("Not found")

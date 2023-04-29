package main

import (
	"database/sql"
	"fmt"
	"os"
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

	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	defer db.Close()

	if conErr != nil {
		return "", conErr
	}

	hash := GenerateSHA256Hash(password)
	token := GenerateSecureToken(tokenDefaultLength)
	_, queryErr := db.Query(
		"insert into users (login, password, token) values (?, ?, ?)",
		login,
		hash,
		token,
	)

	if queryErr != nil {
		return "", queryErr
	}

	return token, nil
}

func (authService) Login(login string, password string) (string, error) {
	if login == "" || password == "" {
		return "", ErrEmpty
	}

	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	defer db.Close()

	if conErr != nil {
		return "", conErr
	}

	hash := GenerateSHA256Hash(password)
	rows, queryErr := db.Query(
		"select * from users where `login` = ? and `password` = ?",
		login,
		hash,
	)

	if queryErr != nil {
		return "", queryErr
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
		_, queryErr := db.Query(
			"update users set token = ? where `login` = ? and `password` = ?",
			token,
			login,
			hash,
		)

		if queryErr != nil {
			return "", queryErr
		}

		return token, nil
	}

	return "", ErrNotFound
}

func (authService) Logout(token string) (string, error) {
	if token == "" {
		return "", ErrEmpty
	}

	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	defer db.Close()

	if conErr != nil {
		return "", conErr
	}

	rows, queryErr := db.Query(
		"select * from users where `token` = ?",
		token,
	)
	fmt.Println(token, rows, queryErr)

	if queryErr != nil {
		return "", queryErr
	}

	if rows.Next() != false {
		u := User{}
		scanErr := rows.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

		if scanErr != nil {
			return "", scanErr
		}

		_, queryErr := db.Query("update users set token = null where `token` = ?", token)

		if queryErr != nil {
			return "", queryErr
		}

		return "Logged out", nil
	}

	return "", ErrNotFound
}

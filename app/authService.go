package main

import (
	"database/sql"
)

type AuthService interface {
	Register(string, string, *sql.DB) (string, error)
	Login(string, string, *sql.DB) (string, error)
	Logout(string, *sql.DB) (string, error)
}

type authService struct{}

func (authService) Register(login string, password string, db *sql.DB) (string, error) {
	if login == "" || password == "" {
		return "", ErrEmpty
	}

	hash := GenerateSHA256Hash(password)
	token := GenerateSecureToken(tokenDefaultLength)
	rows, queryErr := db.Query(
		"insert into users (login, password, token) values (?, ?, ?)",
		login,
		hash,
		token,
	)

	if queryErr != nil {
		return "", queryErr
	}

	defer rows.Close()

	return token, nil
}

func (authService) Login(login string, password string, db *sql.DB) (string, error) {
	if login == "" || password == "" {
		return "", ErrEmpty
	}

	hash := GenerateSHA256Hash(password)
	authUser, authQueryErr := db.Query(
		"select * from users where `login` = ? and `password` = ?",
		login,
		hash,
	)

	if authQueryErr != nil {
		return "", authQueryErr
	}

	defer authUser.Close()

	u := User{}
	if authUser.Next() {
		authUserScanErr := authUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

		if authUserScanErr != nil {
			return "", ErrNotFound
		}

		if u.Token != nil {
			return *u.Token, nil
		}

		token := GenerateSecureToken(tokenDefaultLength)
		rows, queryErr := db.Query(
			"update users set token = ? where `login` = ? and `password` = ?",
			token,
			login,
			hash,
		)

		if queryErr != nil {
			return "", queryErr
		}

		defer rows.Close()

		return token, nil
	}

	return "", ErrNotFound
}

func (authService) Logout(token string, db *sql.DB) (string, error) {
	if token == "" {
		return "", ErrEmpty
	}

	authUser, authQueryErr := db.Query(
		"select * from users where `token` = ?",
		token,
	)

	if authQueryErr != nil {
		return "", authQueryErr
	}

	defer authUser.Close()

	u := User{}
	if authUser.Next() {
		authUserScanErr := authUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

		if authUserScanErr != nil {
			return "", ErrNotFound
		}

		rows, queryErr := db.Query("update users set token = null where `token` = ?", token)

		if queryErr != nil {
			return "", queryErr
		}

		defer rows.Close()

		return "Logged out", nil
	}

	return "", ErrNotFound
}

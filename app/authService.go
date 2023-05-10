package main

import (
	"database/sql"
	"os"
)

type AuthService interface {
	Register(string, string) (string, error)
	Login(string, string) (string, error)
	Logout(string) (string, error)
}

type authService struct{}

func (authService) Register(login string, password string) (string, error) {
	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	if conErr != nil {
		return "", conErr
	}
	defer db.Close()
	if login == "" || password == "" {
		return "", ErrEmpty
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
	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	if conErr != nil {
		return "", conErr
	}
	defer db.Close()
	if login == "" || password == "" {
		return "", ErrEmpty
	}

	hash := GenerateSHA256Hash(password)
	authUser := db.QueryRow(
		"select * from users where `login` = ? and `password` = ?",
		login,
		hash,
	)
	u := User{}
	authUserScanErr := authUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

	if authUserScanErr != nil {
		return "", ErrNotFound
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

func (authService) Logout(token string) (string, error) {
	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	if conErr != nil {
		return "", conErr
	}
	defer db.Close()
	if token == "" {
		return "", ErrEmpty
	}

	authUser := db.QueryRow(
		"select * from users where `token` = ?",
		token,
	)
	u := User{}
	authUserScanErr := authUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

	if authUserScanErr != nil {
		return "", ErrNotFound
	}

	_, queryErr := db.Query("update users set token = null where `token` = ?", token)

	if queryErr != nil {
		return "", queryErr
	}

	return "Logged out", nil
}

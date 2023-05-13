package repositories

import (
	"database/sql"
	"messagingService/app/errors"
	"messagingService/app/helpers"
	"messagingService/app/models"
)

func RegisterUser(login string, password string, token string, db *sql.DB) error {
	hash := helpers.GenerateSHA256Hash(password)
	rows, queryErr := db.Query(
		"insert into users (login, password, token) values (?, ?, ?)",
		login,
		hash,
		token,
	)

	if queryErr != nil {
		return queryErr
	}

	defer rows.Close()

	return nil
}

func GetUserByLogin(login string, db *sql.DB) (models.User, error) {
	user, queryErr := db.Query(
		"select * from users where `login` = ?",
		login,
	)

	if queryErr != nil {
		return models.User{}, errors.ErrNotFound
	}

	defer user.Close()

	if user.Next() {
		u := models.User{}
		scanErr := user.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

		if scanErr != nil {
			return models.User{}, scanErr
		}

		return u, nil
	}

	return models.User{}, errors.ErrNotFound
}

func GetUserByLoginAndPass(login string, password string, db *sql.DB) (models.User, error) {
	hash := helpers.GenerateSHA256Hash(password)
	user, authQueryErr := db.Query(
		"select * from users where `login` = ? and `password` = ?",
		login,
		hash,
	)

	if authQueryErr != nil {
		return models.User{}, errors.ErrNotFound
	}

	defer user.Close()

	if user.Next() {
		u := models.User{}
		scanErr := user.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

		if scanErr != nil {
			return models.User{}, scanErr
		}

		return u, nil
	}

	return models.User{}, errors.ErrNotFound
}

func GetUserByToken(token string, db *sql.DB) (models.User, error) {
	user, authQueryErr := db.Query(
		"select * from users where `token` = ?",
		token,
	)

	if authQueryErr != nil {
		return models.User{}, authQueryErr
	}

	defer user.Close()

	if user.Next() {
		u := models.User{}
		scanErr := user.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

		if scanErr != nil {
			return models.User{}, scanErr
		}

		return u, nil
	}

	return models.User{}, errors.ErrNotFound
}

func LogoutUserByToken(token string, db *sql.DB) error {
	rows, queryErr := db.Query("update users set token = null where `token` = ?", token)

	if queryErr != nil {
		return queryErr
	}

	defer rows.Close()

	return nil
}

func SetUserTokenByLoginAndPass(login string, password string, token string, db *sql.DB) error {
	hash := helpers.GenerateSHA256Hash(password)
	rows, queryErr := db.Query(
		"update users set token = ? where `login` = ? and `password` = ?",
		token,
		login,
		hash,
	)

	if queryErr != nil {
		return queryErr
	}

	defer rows.Close()

	return nil
}

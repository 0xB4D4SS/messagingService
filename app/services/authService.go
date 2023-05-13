package services

import (
	"database/sql"
	"messagingService/app/constants"
	"messagingService/app/errors"
	"messagingService/app/helpers"
	"messagingService/app/repositories"
)

type AuthServiceInterface interface {
	Register(string, string, *sql.DB) (string, error)
	Login(string, string, *sql.DB) (string, error)
	Logout(string, *sql.DB) (string, error)
}

type AuthService struct{}

func (AuthService) Register(login string, password string, db *sql.DB) (string, error) {
	if login == "" || password == "" {
		return "", errors.ErrEmpty
	}

	token := helpers.GenerateSecureToken(constants.TokenDefaultLength)
	registerErr := repositories.RegisterUser(login, password, token, db)

	if registerErr != nil {
		return "", registerErr
	}

	return token, nil
}

func (AuthService) Login(login string, password string, db *sql.DB) (string, error) {
	if login == "" || password == "" {
		return "", errors.ErrEmpty
	}

	authUser, getUserErr := repositories.GetUserByLoginAndPass(login, password, db)

	if getUserErr != nil {
		return "", getUserErr
	}

	if authUser.Id != 0 {
		if authUser.Token != nil {
			return *authUser.Token, nil
		}

		token := helpers.GenerateSecureToken(constants.TokenDefaultLength)
		setTokenErr := repositories.SetUserTokenByLoginAndPass(login, password, token, db)

		if setTokenErr != nil {
			return "", setTokenErr
		}

		return token, nil
	}

	return "", errors.ErrNotFound
}

func (AuthService) Logout(token string, db *sql.DB) (string, error) {
	if token == "" {
		return "", errors.ErrEmpty
	}

	authUser, authUserErr := repositories.GetUserByToken(token, db)

	if authUserErr != nil {
		return "", authUserErr
	}

	if authUser.Id != 0 {
		logoutErr := repositories.LogoutUserByToken(token, db)

		if logoutErr != nil {
			return "", logoutErr
		}

		return "Logged out", nil
	}

	return "", errors.ErrNotFound
}

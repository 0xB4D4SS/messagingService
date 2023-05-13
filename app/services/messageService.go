package services

import (
	"database/sql"
	"messagingService/app/errors"
	"messagingService/app/models"
	"messagingService/app/repositories"
)

type MessageServiceInterface interface {
	Send(string, string, *sql.DB) (string, error)
	Get(string, string, *sql.DB) ([]models.Message, error)
	GetLast(string, string, *sql.DB) (models.Message, error)
}

type MessageService struct{}

func (MessageService) Send(token string, data string, db *sql.DB) (string, error) {
	if token == "" || data == "" {
		return "", errors.ErrEmpty
	}

	authUser, authUserErr := repositories.GetUserByToken(token, db)

	if authUserErr != nil {
		return "", authUserErr
	}

	if authUser.Id != 0 {
		m := models.Message{
			UserId:  authUser.Id,
			Message: data,
		}

		createErr := repositories.CreateMessage(m.UserId, m.Message, db)

		if createErr != nil {
			return "", createErr
		}

		return "ok", nil
	}

	return "", errors.ErrNotFound
}

func (MessageService) Get(token string, login string, db *sql.DB) ([]models.Message, error) {
	if token == "" || login == "" {
		return nil, errors.ErrEmpty
	}

	authUser, authUserErr := repositories.GetUserByToken(token, db)

	if authUserErr != nil {
		return nil, authUserErr
	}

	if authUser.Id != 0 {
		selectUser, selectUserErr := repositories.GetUserByLogin(login, db)

		if selectUserErr != nil {
			return nil, selectUserErr
		}

		if selectUser.Id != 0 {
			messages, getMessagesErr := repositories.GetMessagesByUserId(selectUser.Id, db)

			if getMessagesErr != nil {
				return nil, getMessagesErr
			}

			return messages, nil
		}

		return nil, errors.ErrNotFound
	}

	return nil, errors.ErrNotFound
}

func (MessageService) GetLast(token string, login string, db *sql.DB) (models.Message, error) {
	if token == "" || login == "" {
		return models.Message{}, errors.ErrEmpty
	}

	authUser, getAuthUserErr := repositories.GetUserByToken(token, db)

	if getAuthUserErr != nil {
		return models.Message{}, getAuthUserErr
	}

	if authUser.Id != 0 {
		selectUser, selectUserErr := repositories.GetUserByLogin(login, db)

		if selectUserErr != nil {
			return models.Message{}, selectUserErr
		}

		if selectUser.Id != 0 {
			message, latestMessageErr := repositories.GetLatestMessageByUserId(selectUser.Id, db)

			if latestMessageErr != nil {
				return models.Message{}, latestMessageErr
			}

			return message, nil
		}

		return models.Message{}, errors.ErrNotFound
	}

	return models.Message{}, errors.ErrNotFound
}

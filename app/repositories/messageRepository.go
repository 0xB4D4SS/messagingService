package repositories

import (
	"database/sql"
	"messagingService/app/errors"
	"messagingService/app/models"
)

func CreateMessage(userId int, message string, db *sql.DB) error {
	rows, queryErr := db.Query("insert into messages (user_id, data) values (?, ?)", userId, message)

	if queryErr != nil {
		return errors.ErrCouldNotInsert
	}

	defer rows.Close()

	return nil
}

func GetMessagesByUserId(userId int, db *sql.DB) ([]models.Message, error) {
	messages, queryErr := db.Query("select * from messages where user_id = ?", userId)

	if queryErr != nil {
		return nil, errors.ErrNotFound
	}

	defer messages.Close()

	var data []models.Message
	for messages.Next() {
		m := models.Message{}
		scanErr := messages.Scan(&m.Id, &m.UserId, &m.Message)
		if scanErr != nil {
			return nil, scanErr
		}

		data = append(data, m)
	}

	return data, nil
}

func GetLatestMessageByUserId(userId int, db *sql.DB) (models.Message, error) {
	message, messageQueryErr := db.Query("select * from messages where user_id = ? order by id desc limit 1", userId)

	if messageQueryErr != nil {
		return models.Message{}, errors.ErrNotFound
	}

	defer message.Close()

	if message.Next() {
		m := models.Message{}
		messageScanErr := message.Scan(&m.Id, &m.UserId, &m.Message)

		if messageScanErr != nil {
			return models.Message{}, errors.ErrNotFound
		}

		return m, nil
	}

	return models.Message{}, errors.ErrNotFound
}

package main

import (
	"database/sql"
)

type MessageService interface {
	Send(string, string, *sql.DB) (string, error)
	Get(string, string, *sql.DB) ([]Message, error)
	GetLast(string, string, *sql.DB) (*Message, error)
}

type messageService struct{}

func (messageService) Send(token string, data string, db *sql.DB) (string, error) {
	if token == "" || data == "" {
		return "", ErrEmpty
	}

	authUser, authQueryErr := db.Query("select * from users where token = ?", token)

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

		m := Message{
			UserId:  u.Id,
			Message: data,
		}

		rows, queryErr := db.Query("insert into messages (user_id, data) values (?, ?)", m.UserId, m.Message)

		if queryErr != nil {
			return "", queryErr
		}

		defer rows.Close()

		return "ok", nil
	}

	return "", ErrNotFound
}

func (messageService) Get(token string, login string, db *sql.DB) ([]Message, error) {
	if token == "" || login == "" {
		return nil, ErrEmpty
	}

	authU := User{}
	authUser, authQueryErr := db.Query("select * from users where token = ?", token)

	if authQueryErr != nil {
		return nil, authQueryErr
	}

	defer authUser.Close()

	if authUser.Next() {
		authUserScanErr := authUser.Scan(&authU.Id, &authU.Login, &authU.Password, &authU.Token)

		if authUserScanErr != nil {
			return nil, ErrNotFound
		}

		selectUser, selectQueryErr := db.Query("select * from users where login = ?", login)

		if selectQueryErr != nil {
			return nil, selectQueryErr
		}

		defer selectUser.Close()

		u := User{}

		if selectUser.Next() {
			selectUserScanErr := selectUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

			if selectUserScanErr != nil {
				return nil, ErrNotFound
			}

			messages, queryErr := db.Query("select * from messages where user_id = ?", u.Id)

			if queryErr != nil {
				return nil, queryErr
			}

			defer messages.Close()

			var data []Message
			for messages.Next() {
				m := Message{}
				scanErr := messages.Scan(&m.Id, &m.UserId, &m.Message)
				if scanErr != nil {
					return nil, scanErr
				}

				data = append(data, m)
			}

			return data, nil
		}

		return nil, ErrNotFound
	}

	return nil, ErrNotFound
}

func (messageService) GetLast(token string, login string, db *sql.DB) (*Message, error) {
	if token == "" || login == "" {
		return nil, ErrEmpty
	}

	authU := User{}
	authUser, authQueryErr := db.Query("select * from users where token = ?", token)

	if authQueryErr != nil {
		return nil, authQueryErr
	}

	defer authUser.Close()

	if authUser.Next() {
		authUserScanErr := authUser.Scan(&authU.Id, &authU.Login, &authU.Password, &authU.Token)

		if authUserScanErr != nil {
			return nil, ErrNotFound
		}

		selectUser, selectQueryErr := db.Query("select * from users where login = ?", login)

		if selectQueryErr != nil {
			return nil, selectQueryErr
		}

		defer selectUser.Close()

		if selectUser.Next() {
			u := User{}
			userScanErr := selectUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

			if userScanErr != nil {
				return nil, ErrNotFound
			}

			message, messageQueryErr := db.Query("select * from messages where user_id = ? order by id desc limit 1", u.Id)

			if messageQueryErr != nil {
				return nil, messageQueryErr
			}

			defer message.Close()

			if message.Next() {
				m := Message{}
				messageScanErr := message.Scan(&m.Id, &m.UserId, &m.Message)

				if messageScanErr != nil {
					return nil, ErrNotFound
				}

				return &m, nil
			}

			return nil, ErrNotFound
		}

		return nil, ErrNotFound
	}

	return nil, ErrNotFound
}

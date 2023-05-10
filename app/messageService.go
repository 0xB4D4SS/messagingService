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

	authUser := db.QueryRow("select * from users where token = ?", token)
	u := User{}
	authUserScanErr := authUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

	if authUserScanErr != nil {
		return "", ErrNotFound
	}

	m := Message{
		UserId:  u.Id,
		Message: data,
	}

	_, queryErr := db.Query("insert into messages (user_id, data) values (?, ?)", m.UserId, m.Message)

	if queryErr != nil {
		return "", queryErr
	}

	return "ok", nil
}

func (messageService) Get(token string, login string, db *sql.DB) ([]Message, error) {
	if token == "" || login == "" {
		return nil, ErrEmpty
	}

	authU := User{}
	authUser := db.QueryRow("select * from users where token = ?", token)
	authUserScanErr := authUser.Scan(&authU.Id, &authU.Login, &authU.Password, &authU.Token)

	if authUserScanErr != nil {
		return nil, ErrNotFound
	}

	selectUser := db.QueryRow("select * from users where login = ?", login)
	u := User{}
	selectUserScanErr := selectUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

	if selectUserScanErr != nil {
		return nil, ErrNotFound
	}

	messages, queryErr := db.Query("select * from messages where user_id = ?", u.Id)

	if queryErr != nil {
		return nil, queryErr
	}

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

func (messageService) GetLast(token string, login string, db *sql.DB) (*Message, error) {
	if token == "" || login == "" {
		return nil, ErrEmpty
	}

	authU := User{}
	authUser := db.QueryRow("select * from users where token = ?", token)
	authUserScanErr := authUser.Scan(&authU.Id, &authU.Login, &authU.Password, &authU.Token)

	if authUserScanErr != nil {
		return nil, ErrNotFound
	}

	selectUser := db.QueryRow("select * from users where login = ?", login)
	u := User{}
	userScanErr := selectUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

	if userScanErr != nil {
		return nil, ErrNotFound
	}

	message := db.QueryRow("select * from messages where user_id = ? order by id desc limit 1", u.Id)
	m := Message{}
	messageScanErr := message.Scan(&m.Id, &m.UserId, &m.Message)

	if messageScanErr != nil {
		return nil, ErrNotFound
	}

	return &m, nil

}

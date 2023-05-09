package main

import (
	"database/sql"
	"os"
)

type MessageService interface {
	Send(string, string) (string, error)
	Get(string, string) ([]Message, error)
	GetLast(string, string) (*Message, error)
}

type messageService struct{}

func (messageService) Send(token string, data string) (string, error) {
	if token == "" || data == "" {
		return "", ErrEmpty
	}

	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	defer db.Close()

	if conErr != nil {
		return "", conErr
	}

	rows, queryErr := db.Query("select * from users where token = ?", token)

	if queryErr != nil {
		return "", queryErr
	}

	if rows.Next() != false {
		u := User{}
		scanErr := rows.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

		if scanErr != nil {
			return "", scanErr
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

	return "", ErrNotFound
}

func (messageService) Get(token string, login string) ([]Message, error) {
	if token == "" || login == "" {
		return nil, ErrEmpty
	}

	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	defer db.Close()

	if conErr != nil {
		return nil, conErr
	}

	userRow, queryErr := db.Query("select * from users where token = ?", token)

	if queryErr != nil {
		return nil, queryErr
	}

	if userRow.Next() != false {
		selectUser, queryErr := db.Query("select * from users where login = ?", login)

		if queryErr != nil {
			return nil, queryErr
		}

		if selectUser.Next() != false {
			u := User{}
			scanErr := selectUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

			if scanErr != nil {
				return nil, scanErr
			}

			rows, queryErr := db.Query("select * from messages where user_id = ?", u.Id)

			if queryErr != nil {
				return nil, queryErr
			}

			var data []Message
			for rows.Next() {
				m := Message{}
				scanErr := rows.Scan(&m.Id, &m.UserId, &m.Message)
				if scanErr != nil {
					return nil, scanErr
				}

				data = append(data, m)
			}
			return data, nil
		}
		return nil, ErrNotFound
	}
	return nil, nil
}

func (messageService) GetLast(token string, login string) (*Message, error) {
	if token == "" || login == "" {
		return nil, ErrEmpty
	}

	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	defer db.Close()

	if conErr != nil {
		return nil, conErr
	}

	userRow, queryErr := db.Query("select * from users where token = ?", token)

	if queryErr != nil {
		return nil, queryErr
	}

	if userRow.Next() != false {
		selectUser, queryErr := db.Query("select * from users where login = ?", login)

		if queryErr != nil {
			return nil, queryErr
		}

		if selectUser.Next() != false {
			u := User{}
			scanErr := selectUser.Scan(&u.Id, &u.Login, &u.Password, &u.Token)

			if scanErr != nil {
				return nil, scanErr
			}

			rows, queryErr := db.Query("select * from messages where user_id = ? order by id desc limit 1", u.Id)

			if queryErr != nil {
				return nil, queryErr
			}

			m := Message{}
			for rows.Next() {
				scanErr := rows.Scan(&m.Id, &m.UserId, &m.Message)
				if scanErr != nil {
					return nil, scanErr
				}
			}

			if m.Id == 0 {
				return &m, ErrNotFound
			}

			return &m, nil
		}
		return nil, ErrNotFound
	}
	return nil, nil
}

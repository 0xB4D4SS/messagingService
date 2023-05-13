package models

type User struct {
	Id       int
	Login    string
	Password string
	Token    *string
}

type Message struct {
	Id      int
	UserId  int
	Message string
}

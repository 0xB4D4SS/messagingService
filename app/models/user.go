package models

type User struct {
	Id       int
	Login    string
	Password string
	Token    *string
}
package responses

import "messagingService/app/models"

type RegisterResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LogoutResponse struct {
	Response string `json:"response"`
}

type SendResponse struct {
	Response string `json:"response"`
}

type GetResponse struct {
	Data []models.Message `json:"data"`
}

type GetLastResponse struct {
	Data models.Message `json:"data"`
}

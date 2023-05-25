package responses

import "messagingService/app/models"

type SendResponse struct {
	Response string `json:"response"`
}

type GetResponse struct {
	Data []models.Message `json:"data"`
}

type GetLastResponse struct {
	Data models.Message `json:"data"`
}

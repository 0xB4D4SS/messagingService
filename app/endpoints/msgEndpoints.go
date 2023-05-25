package endpoints

import (
	"database/sql"
	"github.com/go-kit/kit/endpoint"
	"messagingService/app/services"
)

type MessageEndpoints struct {
	SendEndpoint    endpoint.Endpoint
	GetEndpoint     endpoint.Endpoint
	GetLastEndpoint endpoint.Endpoint
}

func MakeMessageEndpoints(msgSvc services.MessageServiceInterface, db *sql.DB) MessageEndpoints {
	return MessageEndpoints{
		SendEndpoint:    MakeSendEndpoint(msgSvc, db),
		GetEndpoint:     MakeGetEndpoint(msgSvc, db),
		GetLastEndpoint: MakeGetLastEndpoint(msgSvc, db),
	}
}

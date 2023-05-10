package main

import (
	"database/sql"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	RegisterEndpoint endpoint.Endpoint
	LoginEndpoint    endpoint.Endpoint
	LogoutEndpoint   endpoint.Endpoint
	SendEndpoint     endpoint.Endpoint
	GetEndpoint      endpoint.Endpoint
	GetLastEndpoint  endpoint.Endpoint
}

func MakeServerEndpoints(authSvc AuthService, msgSvc MessageService, db *sql.DB) Endpoints {
	return Endpoints{
		RegisterEndpoint: makeRegisterEndpoint(authSvc, db),
		LoginEndpoint:    makeLoginEndpoint(authSvc, db),
		LogoutEndpoint:   makeLogoutEndpoint(authSvc, db),
		SendEndpoint:     makeSendEndpoint(msgSvc, db),
		GetEndpoint:      makeGetEndpoint(msgSvc, db),
		GetLastEndpoint:  makeGetLastEndpoint(msgSvc, db),
	}
}

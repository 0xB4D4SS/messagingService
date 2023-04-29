package main

import "github.com/go-kit/kit/endpoint"

type Endpoints struct {
	RegisterEndpoint endpoint.Endpoint
	LoginEndpoint    endpoint.Endpoint
	LogoutEndpoint   endpoint.Endpoint
	SendEndpoint     endpoint.Endpoint
	GetEndpoint      endpoint.Endpoint
}

func MakeServerEndpoints(authSvc AuthService, msgSvc MessageService) Endpoints {
	return Endpoints{
		RegisterEndpoint: makeRegisterEndpoint(authSvc),
		LoginEndpoint:    makeLoginEndpoint(authSvc),
		LogoutEndpoint:   makeLogoutEndpoint(authSvc),
		SendEndpoint:     makeSendEndpoint(msgSvc),
		GetEndpoint:      makeGetEndpoint(msgSvc),
	}
}

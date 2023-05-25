package endpoints

import (
	"context"
	"database/sql"
	"github.com/go-kit/kit/endpoint"
	"messagingService/app/requests"
	"messagingService/app/responses"
	"messagingService/app/services"
)

func MakeRegisterEndpoint(svc services.AuthServiceInterface, db *sql.DB) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.RegisterRequest)
		token, err := svc.Register(req.Login, req.Password, db)
		return responses.RegisterResponse{Token: token}, err
	}
}

func MakeLoginEndpoint(svc services.AuthServiceInterface, db *sql.DB) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.LoginRequest)
		token, err := svc.Login(req.Login, req.Password, db)
		return responses.LoginResponse{Token: token}, err
	}
}

func MakeLogoutEndpoint(svc services.AuthServiceInterface, db *sql.DB) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.LogoutRequest)
		response, err := svc.Logout(req.Token, db)
		return responses.LogoutResponse{Response: response}, err
	}
}

func MakeSendEndpoint(svc services.MessageServiceInterface, db *sql.DB) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.SendRequest)
		response, err := svc.Send(req.Token, req.Data, db)
		return responses.SendResponse{Response: response}, err
	}
}

func MakeGetEndpoint(svc services.MessageServiceInterface, db *sql.DB) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.GetRequest)
		data, err := svc.Get(req.Token, req.Login, db)
		return responses.GetResponse{Data: data}, err
	}
}

func MakeGetLastEndpoint(svc services.MessageServiceInterface, db *sql.DB) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.GetRequest)
		data, err := svc.GetLast(req.Token, req.Login, db)
		return responses.GetLastResponse{Data: data}, err
	}
}

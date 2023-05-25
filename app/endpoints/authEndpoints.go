package endpoints

import (
	"database/sql"
	"github.com/go-kit/kit/endpoint"
	"messagingService/app/services"
)

type AuthEndpoints struct {
	RegisterEndpoint endpoint.Endpoint
	LoginEndpoint    endpoint.Endpoint
	LogoutEndpoint   endpoint.Endpoint
}

func MakeAuthEndpoints(authSvc services.AuthServiceInterface, db *sql.DB) AuthEndpoints {
	return AuthEndpoints{
		RegisterEndpoint: MakeRegisterEndpoint(authSvc, db),
		LoginEndpoint:    MakeLoginEndpoint(authSvc, db),
		LogoutEndpoint:   MakeLogoutEndpoint(authSvc, db),
	}
}

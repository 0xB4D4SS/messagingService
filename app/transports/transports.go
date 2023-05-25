package transports

import (
	"database/sql"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"messagingService/app/endpoints"
	"messagingService/app/middleware"
	"messagingService/app/services"
	"net/http"
)

func MakeHTTPHandler(authSvc services.AuthServiceInterface, msgSvc services.MessageServiceInterface, db *sql.DB) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.SetContentType)
	authEndpoints := endpoints.MakeAuthEndpoints(authSvc, db)
	msgEndpoints := endpoints.MakeMessageEndpoints(msgSvc, db)

	r.Methods("POST").Path("/register").Handler(
		httptransport.NewServer(
			authEndpoints.RegisterEndpoint,
			decodeRegisterRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/login").Handler(
		httptransport.NewServer(
			authEndpoints.LoginEndpoint,
			decodeLoginRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/logout").Handler(
		httptransport.NewServer(
			authEndpoints.LogoutEndpoint,
			decodeLogoutRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/send").Handler(
		httptransport.NewServer(
			msgEndpoints.SendEndpoint,
			decodeSendRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/get").Handler(
		httptransport.NewServer(
			msgEndpoints.GetEndpoint,
			decodeGetRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/get-last").Handler(
		httptransport.NewServer(
			msgEndpoints.GetLastEndpoint,
			decodeGetLastRequest,
			encodeResponse,
		))

	return r
}

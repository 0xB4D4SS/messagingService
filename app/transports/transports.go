package transports

import (
	"context"
	"database/sql"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"messagingService/app/endpoints"
	"messagingService/app/requests"
	"messagingService/app/services"
	"net/http"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func MakeHTTPHandler(authSvc services.AuthServiceInterface, msgSvc services.MessageServiceInterface, db *sql.DB) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	e := endpoints.MakeServerEndpoints(authSvc, msgSvc, db)

	r.Methods("POST").Path("/register").Handler(
		httptransport.NewServer(
			e.RegisterEndpoint,
			decodeRegisterRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/login").Handler(
		httptransport.NewServer(
			e.LoginEndpoint,
			decodeLoginRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/logout").Handler(
		httptransport.NewServer(
			e.LogoutEndpoint,
			decodeLogoutRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/send").Handler(
		httptransport.NewServer(
			e.SendEndpoint,
			decodeSendRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/get").Handler(
		httptransport.NewServer(
			e.GetEndpoint,
			decodeGetRequest,
			encodeResponse,
		))
	r.Methods("POST").Path("/get-last").Handler(
		httptransport.NewServer(
			e.GetLastEndpoint,
			decodeGetLastRequest,
			encodeResponse,
		))

	return r
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeLogoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.LogoutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSendRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.SendRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.GetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetLastRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.GetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

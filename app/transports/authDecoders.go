package transports

import (
	"context"
	"encoding/json"
	"messagingService/app/requests"
	"net/http"
)

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

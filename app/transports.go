package main

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type registerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type registerResponse struct {
	Token string `json:"token"`
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

type logoutRequest struct {
	Token string `json:"token"`
}

type logoutResponse struct {
	Response string `json:"response"`
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request registerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request loginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeLogoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request logoutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)

func makeRegisterEndpoint(svc AuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(registerRequest)
		token, err := svc.Register(req.Login, req.Password)
		return registerResponse{token}, err
	}
}

func makeLoginEndpoint(svc AuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		token, err := svc.Login(req.Login, req.Password)
		return loginResponse{token}, err
	}
}

func makeLogoutEndpoint(svc AuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(logoutRequest)
		response, err := svc.Logout(req.Token)
		return logoutResponse{response}, err
	}
}

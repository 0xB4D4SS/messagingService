package main

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler(authSvc AuthService, msgSvc MessageService) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(authSvc, msgSvc)

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

func decodeSendRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sendRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetLastRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getRequest
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

func makeSendEndpoint(svc MessageService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(sendRequest)
		response, err := svc.Send(req.Token, req.Data)
		return sendResponse{response}, err
	}
}

func makeGetEndpoint(svc MessageService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		data, err := svc.Get(req.Token, req.Login)
		return getResponse{data}, err
	}
}

func makeGetLastEndpoint(svc MessageService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		data, err := svc.GetLast(req.Token, req.Login)
		return getLastResponse{*data}, err
	}
}

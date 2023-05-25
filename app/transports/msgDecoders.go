package transports

import (
	"context"
	"encoding/json"
	"messagingService/app/requests"
	"net/http"
)

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

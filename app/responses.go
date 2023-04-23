package main

type registerResponse struct {
	Token string `json:"token"`
}

type loginResponse struct {
	Token string `json:"token"`
}

type logoutResponse struct {
	Response string `json:"response"`
}

type sendResponse struct {
	Response string `json:"response"`
}

type getResponse struct {
	Data []Message `json:"data"`
}

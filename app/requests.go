package main

type registerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type logoutRequest struct {
	Token string `json:"token"`
}

type sendRequest struct {
	Token string `json:"token"`
	Data  string `json:"data"`
}

type getRequest struct {
	Token string `json:"token"`
	Login string `json:"login"`
}

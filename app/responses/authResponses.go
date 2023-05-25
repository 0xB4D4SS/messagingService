package responses

type RegisterResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LogoutResponse struct {
	Response string `json:"response"`
}

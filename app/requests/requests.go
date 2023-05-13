package requests

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	Token string `json:"token"`
}

type SendRequest struct {
	Token string `json:"token"`
	Data  string `json:"data"`
}

type GetRequest struct {
	Token string `json:"token"`
	Login string `json:"login"`
}

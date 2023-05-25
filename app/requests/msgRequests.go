package requests

type SendRequest struct {
	Token string `json:"token"`
	Data  string `json:"data"`
}

type GetRequest struct {
	Token string `json:"token"`
	Login string `json:"login"`
}

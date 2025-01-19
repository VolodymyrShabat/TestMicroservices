package models

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type IsTokenValidRequest struct {
	Token string `json:"token"`
}

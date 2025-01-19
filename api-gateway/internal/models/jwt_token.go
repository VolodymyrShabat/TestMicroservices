package models

type JWTTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type IsTokenValid struct {
	IsTokenValid bool `json:"isTokenValid"`
}

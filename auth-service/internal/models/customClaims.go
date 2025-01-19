package models

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	Username string `json:"userId"`
	jwt.StandardClaims
}

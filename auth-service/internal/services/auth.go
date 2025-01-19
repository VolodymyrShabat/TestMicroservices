package services

import (
	"fmt"
	"github.com/VolodymyrShabat/TestMicroservices/auth-service/internal/models"
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthService struct {
	JwtSecretKey []byte
	Salt         []byte
}

func NewAuthService(JwtSecretKey, Salt []byte) *AuthService {
	return &AuthService{
		JwtSecretKey: JwtSecretKey,
		Salt:         Salt,
	}
}

func (as *AuthService) CreateJWT(username string, isReset bool) (response *models.Token, err error) {
	accessClaims := &models.CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
	}
	if isReset {
		accessClaims.StandardClaims.ExpiresAt = time.Now().Add(10 * time.Minute).Unix()
	} else {
		accessClaims.StandardClaims.ExpiresAt = time.Now().Add(30 * time.Minute).Unix()
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(as.JwtSecretKey)
	if err != nil {
		return nil, err
	}

	refreshClaims := &models.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(8 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(as.JwtSecretKey)
	if err != nil {
		return nil, err
	}

	return &models.Token{AccessToken: accessToken, RefreshToken: refreshToken}, err
}

func (as *AuthService) VerifyToken(signedToken string) (models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &models.CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return 0, fmt.Errorf("unexpected signing method")
			}
			return as.JwtSecretKey, nil
		})
	if err != nil {
		return models.CustomClaims{}, fmt.Errorf("parse token: %w", err)
	}

	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok {
		return models.CustomClaims{}, fmt.Errorf("parse claims: %w", err)
	}

	return models.CustomClaims{
		Username: claims.Username,
	}, nil
}

//func (as *AuthService) RefreshToken(refreshToken string) (*models.Token, error) {
//	claims, err := as.VerifyToken(refreshToken)
//	if err != nil {
//		return nil, fmt.Errorf("verify token: %w", err)
//	}
//
//	user, err := GetUserById(claims.Username)
//	if err != nil {
//		return nil, fmt.Errorf("find user: %w", err)
//	}
//
//	tokenPair, err := as.CreateJWT(user.Username, false)
//	if err != nil {
//		return nil, fmt.Errorf("create token pair: %w", err)
//	}
//
//	return tokenPair, nil
//}

package services

import (
	"context"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/models"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/utils"
	authpb "github.com/VolodymyrShabat/TestMicroservices/auth-service/pkg/proto"
	"time"
)

type AuthService struct {
	AuthClient authpb.AuthServiceClient
	Convertor  *utils.Convertor
}

func NewAuthService(authClient authpb.AuthServiceClient) *AuthService {
	return &AuthService{
		AuthClient: authClient,
		Convertor:  utils.NewConvertor(),
	}
}

func (as *AuthService) SignIn(username string, password string) (*models.JWTTokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := as.AuthClient.Login(ctx, &authpb.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return as.Convertor.ConvertToken(res), nil
}

func (as *AuthService) ValidateToken(token string) (*models.IsTokenValid, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := as.AuthClient.ValidateToken(ctx, &authpb.ValidateTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	return as.Convertor.ConvertIsValid(res), nil
}

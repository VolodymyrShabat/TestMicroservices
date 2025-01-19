package utils

import (
	"github.com/VolodymyrShabat/TestMicroservices/auth-service/internal/models"
	authpb "github.com/VolodymyrShabat/TestMicroservices/auth-service/pkg/proto"
)

type Convertor struct {
}

func NewConvertor() *Convertor {
	return &Convertor{}
}

func (c *Convertor) ValidateTokenToProto(IsValid bool) *authpb.ValidateTokenResponse {
	return &authpb.ValidateTokenResponse{
		IsValid: IsValid,
	}
}

func (c *Convertor) JWTTokenToProto(jwtToken *models.Token) *authpb.LoginResponse {
	return &authpb.LoginResponse{
		AccessToken:  jwtToken.AccessToken,
		RefreshToken: jwtToken.RefreshToken,
	}
}

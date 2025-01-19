package utils

import (
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/models"
	authpb "github.com/VolodymyrShabat/TestMicroservices/auth-service/pkg/proto"
	resourcespb "github.com/VolodymyrShabat/TestMicroservices/resource-service/pkg/proto"
)

type Convertor struct {
}

func NewConvertor() *Convertor {
	return &Convertor{}
}

func (converter *Convertor) ConvertBooks(b *resourcespb.Books) []*models.Book {
	books := make([]*models.Book, len(b.Books))
	for i := range b.Books {
		books[i] = &models.Book{
			Title:  b.Books[i].Title,
			Author: b.Books[i].Author,
		}
	}
	return books
}

func (converter *Convertor) ConvertUsers(u *resourcespb.Users) []*models.User {
	users := make([]*models.User, len(u.Users))
	for i := range u.Users {
		users[i] = &models.User{
			Username: u.Users[i].Username,
			Email:    u.Users[i].Email,
			Roles:    u.Users[i].Roles,
		}
	}
	return users
}

func (converter *Convertor) ConvertToken(b *authpb.LoginResponse) *models.JWTTokenResponse {
	return &models.JWTTokenResponse{
		AccessToken:  b.AccessToken,
		RefreshToken: b.RefreshToken,
	}
}

func (converter *Convertor) ConvertIsValid(b *authpb.ValidateTokenResponse) *models.IsTokenValid {
	return &models.IsTokenValid{
		IsTokenValid: b.IsValid,
	}
}

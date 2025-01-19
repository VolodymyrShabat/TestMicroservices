package utils

import (
	"fmt"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/models"
	resourcepb "github.com/VolodymyrShabat/TestMicroservices/resource-service/pkg/proto"
)

type Convertor struct {
}

func NewConvertor() *Convertor {
	return &Convertor{}
}

func (c *Convertor) BookConvertToProto(book []*models.Book) *resourcepb.Books {
	books := &resourcepb.Books{}
	b := make([]*resourcepb.Book, len(book))
	for i, v := range book {
		b[i] = &resourcepb.Book{
			Id:     fmt.Sprint(v.Id),
			Author: v.Author,
			Title:  v.Title,
		}
	}
	books.Books = b
	return books
}

func (c *Convertor) UserConvertToProto(book []*models.User) *resourcepb.Users {
	users := &resourcepb.Users{}
	u := make([]*resourcepb.User, len(book))
	for i, v := range book {
		u[i] = &resourcepb.User{
			Id:       fmt.Sprint(v.Id),
			Username: v.Username,
			Email:    v.Email,
			Roles:    v.Roles,
		}
	}
	users.Users = u
	return users
}

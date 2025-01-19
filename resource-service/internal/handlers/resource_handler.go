package handlers

import (
	"context"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/services"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/utils"
	resourcepb "github.com/VolodymyrShabat/TestMicroservices/resource-service/pkg/proto"
)

type ResourceHandler struct {
	resourcepb.UnimplementedResourceServiceServer
	UserService *services.UserService
	BookService *services.BookService
	Convertor   *utils.Convertor
}

func NewResourceHandler(userService *services.UserService, bookService *services.BookService) *ResourceHandler {
	return &ResourceHandler{
		UserService: userService,
		BookService: bookService,
		Convertor:   utils.NewConvertor(),
	}
}

func (h *ResourceHandler) GetUsers(context.Context, *resourcepb.EmptyRequest) (*resourcepb.Users, error) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		return nil, err
	}
	return h.Convertor.UserConvertToProto(users), nil
}

func (h *ResourceHandler) GetBooks(context.Context, *resourcepb.EmptyRequest) (*resourcepb.Books, error) {
	books, err := h.BookService.GetBooks()
	if err != nil {
		return nil, err
	}
	return h.Convertor.BookConvertToProto(books), nil
}

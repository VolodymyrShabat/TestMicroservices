package handlers

import (
	"context"
	resourcepb "github.com/VolodymyrShabat/TestMicroservices/resource-service/pkg/proto"
)

type ResourceHandler struct {
	resourcepb.UnimplementedResourceServiceServer
}

func (h *ResourceHandler) GetUsers(context.Context, *resourcepb.EmptyRequest) (*resourcepb.Users, error) {
	// Your business logic (e.g., verify credentia
	return &resourcepb.Users{}, nil
}

func (h *ResourceHandler) GetBooks(context.Context, *resourcepb.EmptyRequest) (*resourcepb.Books, error) {
	// Your token validation logic
	return &resourcepb.Books{}, nil
}

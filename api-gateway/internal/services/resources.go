package services

import (
	"context"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/models"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/utils"
	resourcespb "github.com/VolodymyrShabat/TestMicroservices/resource-service/pkg/proto"
	"time"
)

type ResourcesService struct {
	ResourcesClient resourcespb.ResourceServiceClient
	Convertor       *utils.Convertor
}

func NewResourcesService(authClient resourcespb.ResourceServiceClient) *ResourcesService {
	return &ResourcesService{
		ResourcesClient: authClient,
		Convertor:       utils.NewConvertor(),
	}
}

func (r ResourcesService) GetBooks() ([]*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := r.ResourcesClient.GetBooks(ctx, &resourcespb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	return r.Convertor.ConvertBooks(resp), nil
}

func (r ResourcesService) GetUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := r.ResourcesClient.GetUsers(ctx, &resourcespb.EmptyRequest{})
	if err != nil {
		// You can map different error types to different status codes if you like
		return nil, err
	}
	return r.Convertor.ConvertUsers(resp), nil
}

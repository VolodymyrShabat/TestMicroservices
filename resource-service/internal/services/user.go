package services

import (
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/models"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/repository"
)

type UserService struct {
	*repository.UserRepository
}

func NewUserService() *UserService {
	repo := repository.NewUserRepository()
	return &UserService{repo}
}

func (u *UserService) GetUsers() ([]*models.User, error) {
	books, err := u.UserRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return books, nil
}

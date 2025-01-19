package repository

import "github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/models"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

var users = []*models.User{
	{1, "kyrkela", "volodymyrshabat@gmail.com", []string{"user", "vip"}},
	{2, "avbyte", "nazardubenko2@gmail.com", []string{"user"}},
	{3, "yulianam", "yulianamal@gmail.com", []string{"user", "admin"}},
}

func (u *UserRepository) GetUsers() ([]*models.User, error) {
	return users, nil
}

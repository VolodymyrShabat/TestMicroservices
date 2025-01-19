package services

import (
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/models"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/repository"
)

type BookService struct {
	BookRepository *repository.BookRepository
}

func NewBookService() *BookService {
	repo := repository.NewBookRepository()
	return &BookService{BookRepository: repo}
}

func (u *BookService) GetBooks() ([]*models.Book, error) {
	books, err := u.BookRepository.GetBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

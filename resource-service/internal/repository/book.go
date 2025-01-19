package repository

import "github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/models"

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

var books = []*models.Book{
	{1, "Harry Potter 1", "J.K.Rowling"},
	{2, "Harry Potter 2", "J.K.Rowling"},
	{3, "Lord of the Rings", "J.R.R. Tolkien"},
}

func (u *BookRepository) GetBooks() ([]*models.Book, error) {
	return books, nil
}

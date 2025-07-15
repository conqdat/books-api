package service

import (
	"context"
	"fmt"

	"github.com/conqdat/books-api/internal/models"
	repository "github.com/conqdat/books-api/internal/repository/postgres"
)

type BookService struct {
	bookRepo *repository.BookRepository
}

func NewBookService(bookRepo *repository.BookRepository) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	books, err := s.bookRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get books: %w", err)
	}

	return books, nil
}
package books

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/configs"
	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
)

type bookRepository interface {
	AddBook(ctx context.Context, model books.BookModel) error
	DeleteBook(ctx context.Context, title, author string) error
	GetAllBook(ctx context.Context, limit, offset int) (books.GetAllBookResponse, error)
}

type service struct {
	cfg      *configs.Config
	bookRepo bookRepository
}

func NewService(cfg *configs.Config, bookrepo bookRepository) *service {
	return &service{
		cfg:      cfg,
		bookRepo: bookrepo,
	}
}

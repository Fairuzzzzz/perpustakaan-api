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
	GetBookByID(ctx context.Context, id int64) (*books.BookModel, error)
	UpdateBook(ctx context.Context, model books.BookModel) error
	GetBookByTitleAndAuthor(ctx context.Context, title, author string) (*books.BookModel, error)
	BorrowBook(ctx context.Context, model books.BorrowModel) error
	DecrementAvailableCopies(ctx context.Context, bookID int64) error
	ReturnBook(ctx context.Context, userID, bookID int64) error
	IncrementAvailableCopies(ctx context.Context, bookID int64) error
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

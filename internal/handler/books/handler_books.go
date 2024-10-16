package books

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/middleware"
	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/gin-gonic/gin"
)

type bookService interface {
	AddBook(ctx context.Context, req books.AddBookRequest) error
	DeleteBook(ctx context.Context, req books.DeleteBookRequest) error
	GetAllBook(ctx context.Context, pageSize, pageIndex int) (books.GetAllBookResponse, error)
	UpdateBook(ctx context.Context, req books.UpdateBookRequest) error
	BorrowBook(ctx context.Context, req books.BorrowBookRequest) error
	ReturnBook(ctx context.Context, req books.ReturnBookRequest) error
	GetAllBorrowedBook(ctx context.Context, pageSize, pageIndex int) (books.GetAllBorrowedBookResponse, error)
}

type Handler struct {
	*gin.Engine

	bookSvc bookService
}

func NewHandler(api *gin.Engine, bookSvc bookService) *Handler {
	return &Handler{
		Engine:  api,
		bookSvc: bookSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("books")
	route.Use(middleware.AuthMiddleware())
	route.POST("/borrow-book", h.BorrowBook)
	route.POST("/return-book", h.ReturnBook)
	route.Use(middleware.AdminOnly())
	route.POST("/add-book", h.AddBook)
	route.DELETE("/delete-book", h.DeleteBook)
	route.PUT("/update-book", h.UpdateBook)
	route.GET("/", h.GetAllBook)
	route.GET("/borrowed-book", h.GetAllBorrowedBook)
}

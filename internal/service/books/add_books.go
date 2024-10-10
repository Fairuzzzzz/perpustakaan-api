package books

import (
	"context"
	"strings"
	"time"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/rs/zerolog/log"
)

func (s *service) AddBook(ctx context.Context, req books.AddBookRequest) error {
	bookCategory := strings.Join(req.Category, ",")
	availableCopies := req.TotalCopies
	publicationYear, err := time.Parse("2006-01-02", req.PublicationYear)
	if err != nil {
		log.Error().Err(err).Msg("invalid publication year format")
		return err
	}
	model := books.BookModel{
		Title:           req.Title,
		Author:          req.Author,
		Category:        bookCategory,
		PublicationYear: publicationYear.Format("2006-01-02"),
		TotalCopies:     req.TotalCopies,
		AvailableCopies: availableCopies,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err = s.bookRepo.AddBook(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error add book to repository")
		return err
	}
	return nil
}

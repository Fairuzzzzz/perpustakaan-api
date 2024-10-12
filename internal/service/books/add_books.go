package books

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/rs/zerolog/log"
)

func validatePublicationYear(date string) error {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return errors.New("invalid publication year format, should be in ISO 8601 format")
	}
	return nil
}

func (s *service) AddBook(ctx context.Context, req books.AddBookRequest) error {
	bookCategory := strings.Join(req.Category, ",")
	availableCopies := req.TotalCopies

	if err := validatePublicationYear(req.PublicationYear); err != nil {
		log.Error().Str("year", req.PublicationYear).Msg(err.Error())
		return err
	}

	model := books.BookModel{
		Title:           req.Title,
		Author:          req.Author,
		Category:        bookCategory,
		PublicationYear: req.PublicationYear,
		TotalCopies:     req.TotalCopies,
		AvailableCopies: availableCopies,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := s.bookRepo.AddBook(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error add book to repository")
		return err
	}
	return nil
}

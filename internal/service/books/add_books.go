package books

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/rs/zerolog/log"
)

func validateAndFormatPublicationYear(date string) (string, error) {
	formats := []string{
		"2006-01-02T15:04:05Z07:00", // RFC3339
		"2006-01-02",                // YYYY-MM-DD
		"2006-01-02 15:04:05",       // YYYY-MM-DD HH:MM:SS
		"2006",                      // YYYY
	}

	var parsedTime time.Time
	var err error
	for _, format := range formats {
		parsedTime, err = time.Parse(format, date)
		if err == nil {
			break
		}
	}

	if err != nil {
		return "", errors.New("invalid publication year format")
	}

	return parsedTime.Format("2006-01-02"), nil
}

func (s *service) AddBook(ctx context.Context, req books.AddBookRequest) error {
	bookCategory := strings.Join(req.Category, ",")
	availableCopies := req.TotalCopies

	formattedDate, err := validateAndFormatPublicationYear(req.PublicationYear)
	if err != nil {
		log.Error().Str("year", req.PublicationYear).Msg(err.Error())
		return err
	}

	model := books.BookModel{
		Title:           req.Title,
		Author:          req.Author,
		Category:        bookCategory,
		PublicationYear: formattedDate,
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

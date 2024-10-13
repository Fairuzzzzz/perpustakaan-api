package books

import (
	"context"
	"strings"
	"time"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/rs/zerolog/log"
)

func (s *service) UpdateBook(ctx context.Context, req books.UpdateBookRequest) error {
	existingBook, err := s.bookRepo.GetBookByID(ctx, req.ID)
	if err != nil {
		log.Error().Err(err).Msg("error getting existing book from repository")
		return err
	}

	if req.Title != nil {
		existingBook.Title = *req.Title
	}

	if req.Author != nil {
		existingBook.Author = *req.Author
	}

	if req.Category != nil {
		existingBook.Category = strings.Join(*req.Category, ",")
	}

	if req.PublicationYear != nil {
		formattedDate, err := validateAndFormatPublicationYear(*req.PublicationYear)
		if err != nil {
			log.Error().Str("year", *req.PublicationYear).Msg(err.Error())
			return err
		}
		existingBook.PublicationYear = formattedDate
	}

	if req.TotalCopies != nil {
		existingBook.TotalCopies = *req.TotalCopies
	}

	if req.AvailableCopies != nil {
		existingBook.AvailableCopies = *req.AvailableCopies
	}

	existingBook.UpdatedAt = time.Now()

	err = s.bookRepo.UpdateBook(ctx, *existingBook)
	if err != nil {
		log.Error().Err(err).Msg("error updating book in repository")
		return err
	}
	return nil
}

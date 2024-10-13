package books

import (
	"context"
	"errors"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/rs/zerolog/log"
)

func (s *service) ReturnBook(ctx context.Context, req books.ReturnBookRequest) error {
	book, err := s.bookRepo.GetBookByTitleAndAuthor(ctx, req.Title, req.Author)
	if err != nil {
		log.Error().Err(err).Msg("error getting book from repository")
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	err = s.bookRepo.ReturnBook(ctx, req.UserID, book.ID)
	if err != nil {
		log.Error().Err(err).Msg("error returning book from repository")
		return err
	}

	err = s.bookRepo.IncrementAvailableCopies(ctx, book.ID)
	if err != nil {
		log.Error().Err(err).Msg("error updating book in repository")
		return err
	}
	return nil
}

package books

import (
	"context"
	"errors"
	"time"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/rs/zerolog/log"
)

func (s *service) BorrowBook(ctx context.Context, req books.BorrowBookRequest) error {
	book, err := s.bookRepo.GetBookByTitleAndAuthor(ctx, req.Title, req.Author)
	if err != nil {
		log.Error().Err(err).Msg("error getting book from repository")
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	if book.AvailableCopies <= 0 {
		return errors.New("no book available")
	}

	borrowModel := books.BorrowModel{
		UserID:     req.UserID,
		BookID:     book.ID,
		BorrowDate: time.Now(),
		DueDate:    time.Now().AddDate(0, 0, 7),
		IsReturned: false,
		ReturnDate: nil,
	}

	err = s.bookRepo.BorrowBook(ctx, borrowModel)
	if err != nil {
		log.Error().Err(err).Msg("error borrow book in repository")
		return err
	}

	err = s.bookRepo.DecrementAvailableCopies(ctx, book.ID)
	if err != nil {
		log.Error().Err(err).Msg("error updating book in repository")
		return err
	}
	return nil
}

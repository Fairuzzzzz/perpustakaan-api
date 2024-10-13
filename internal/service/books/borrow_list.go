package books

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllBorrowedBook(ctx context.Context, pageSize, pageIndex int) (books.GetAllBorrowedBookResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)
	response, err := s.bookRepo.GetAllBorrowedBook(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error get all book from database")
		return response, err
	}
	return response, nil
}

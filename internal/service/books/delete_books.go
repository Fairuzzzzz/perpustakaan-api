package books

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/rs/zerolog/log"
)

func (s *service) DeleteBook(ctx context.Context, req books.DeleteBookRequest) error {
	err := s.bookRepo.DeleteBook(ctx, req.Title, req.Author)
	if err != nil {
		log.Error().Err(err).Msg("error deleting books from repository")
		return err
	}
	return nil
}

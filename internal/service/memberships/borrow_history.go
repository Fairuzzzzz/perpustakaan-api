package memberships

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
	"github.com/rs/zerolog/log"
)

func (s *service) GetBorrowHistory(ctx context.Context, userID int64) ([]memberships.BorrowHistory, error) {
	history, err := s.membershipRepo.GetBorrowHistory(ctx, userID)
	if err != nil {
		log.Error().Err(err).Msg("error getting borrow history from repository")
		return nil, err
	}
	return history, err
}

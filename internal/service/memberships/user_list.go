package memberships

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllUser(ctx context.Context, pageSize, pageIndex int) (memberships.GetAllUserResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)
	response, err := s.membershipRepo.GetAllUser(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error get all user from database")
		return response, err
	}
	return response, nil
}

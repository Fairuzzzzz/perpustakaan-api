package memberships

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
	"github.com/rs/zerolog/log"
)

func (s *service) DeleteUsers(ctx context.Context, req memberships.DeleteUserRequest) error {
	err := s.membershipRepo.DeleteUser(ctx, req.Email, req.Username)
	if err != nil {
		log.Error().Err(err).Msg("error deleting users from repository")
		return err
	}
	return nil
}

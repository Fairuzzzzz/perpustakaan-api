package memberships

import (
	"context"
	"errors"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
	"github.com/Fairuzzzzz/perpustakaan-api/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("email not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("email or password is invalid")
		return "", err
	}

	role := user.Role
	userID := user.ID

	// Integrasi fungsi generate JWT
	token, err := jwt.CreateToken(userID, role, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", err
	}

	return token, nil
}

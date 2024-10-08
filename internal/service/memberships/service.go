package memberships

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/configs"
	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
)

type membershipsRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	cfg            *configs.Config
	membershipRepo membershipsRepository
}

func NewService(cfg *configs.Config, membershipsRepo membershipsRepository) *service {
	return &service{
		cfg:            cfg,
		membershipRepo: membershipsRepo,
	}
}

package memberships

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
)

type membershipsRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	membershipRepo membershipsRepository
}

func NewService(membershipsRepo membershipsRepository) *service {
	return &service{
		membershipRepo: membershipsRepo,
	}
}

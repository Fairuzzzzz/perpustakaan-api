package memberships

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/configs"
	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
)

type membershipsRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
	DeleteUser(ctx context.Context, email, username string) error
	GetAllUser(ctx context.Context, limit, offset int) (memberships.GetAllUserResponse, error)
	GetBorrowHistory(ctx context.Context, userID int64) ([]memberships.BorrowHistory, error)
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

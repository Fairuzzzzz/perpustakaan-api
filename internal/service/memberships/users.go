package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("email or username already exists")
	}

	// Hash Password
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Model Database
	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Password:  string(pass),
		Username:  req.Username,
		Role:      "anggota",
		CreatedAt: now,
		UpdatedAt: now,
	}
	return s.membershipRepo.CreateUser(ctx, model)
}

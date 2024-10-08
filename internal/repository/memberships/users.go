package memberships

import (
	"context"
	"database/sql"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `SELECT id, email, password, username, role, created_at, updated_at FROM users WHERE email = ? OR id = ?`

	row := r.db.QueryRowContext(ctx, query, email, username)

	var response memberships.UserModel

	err := row.Scan(&response.ID, &response.Email, &response.Password, &response.Username, &response.Role, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `INSERT INTO users (email, password, username, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username, model.Role, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

package memberships

import (
	"context"
	"database/sql"
	"time"

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

func (r *repository) DeleteUser(ctx context.Context, email, username string) error {
	query := `DELETE FROM users WHERE email = ? AND username = ?`
	_, err := r.db.ExecContext(ctx, query, email, username)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllUser(ctx context.Context, limit, offset int) (memberships.GetAllUserResponse, error) {
	query := `SELECT u.id, u.email, u.username, u.role FROM users u ORDER BY u.updated_at DESC LIMIT ? OFFSET ?`

	var response memberships.GetAllUserResponse
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	data := make([]memberships.User, 0)
	for rows.Next() {
		var model memberships.UserModel
		err := rows.Scan(&model.ID, &model.Email, &model.Username, &model.Role)
		if err != nil {
			return response, err
		}
		data = append(data, memberships.User{
			ID:       model.ID,
			Email:    model.Email,
			Username: model.Username,
			Role:     model.Role,
		})
	}
	response.Data = data
	response.Pagination = memberships.Pagination{
		Limit:  limit,
		Offset: offset,
	}
	return response, nil
}

func (r *repository) GetBorrowHistory(ctx context.Context, userID int64) ([]memberships.BorrowHistory, error) {
	query := `SELECT u.username, b.id, b.title, b.author, br.borrow_date, br.due_date, br.return_date, br.is_returned FROM books b JOIN borrows br ON br.book_id = b.id JOIN users u ON br.user_id = u.id WHERE br.user_id = ? ORDER BY br.borrow_date DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []memberships.BorrowHistory
	for rows.Next() {
		var borrow memberships.BorrowHistory
		var returnDate *time.Time

		err := rows.Scan(&borrow.Username, &borrow.BookID, &borrow.Title, &borrow.Author, &borrow.BorrowDate, &borrow.DueDate, &returnDate, &borrow.IsReturned)
		if err != nil {
			return nil, err
		}

		borrow.ReturnDate = returnDate
		history = append(history, borrow)
	}
	return history, nil
}

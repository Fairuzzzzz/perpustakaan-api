package books

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
)

func (r *repository) AddBook(ctx context.Context, model books.BookModel) error {
	model.AvailableCopies = model.TotalCopies

	query := `INSERT INTO books (title, author, category, publication_year, total_copies, available_copies, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.Title, model.Author, model.Category, model.PublicationYear, model.TotalCopies, model.AvailableCopies, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

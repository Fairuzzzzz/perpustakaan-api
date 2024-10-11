package books

import (
	"context"
	"strings"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
)

func (r *repository) AddBook(ctx context.Context, model books.BookModel) error {
	query := `INSERT INTO books (title, author, category, publication_year, total_copies, available_copies, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.Title, model.Author, model.Category, model.PublicationYear, model.TotalCopies, model.AvailableCopies, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteBook(ctx context.Context, title, author string) error {
	query := `DELETE FROM books WHERE title = ? AND author = ?`
	_, err := r.db.ExecContext(ctx, query, title, author)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllBook(ctx context.Context, limit, offset int) (books.GetAllBookResponse, error) {
	query := `SELECT b.id, b.title, b.author, b.category, b.publication_year, b.total_copies, b.available_copies FROM books b ORDER BY b.updated_at DESC LIMIT ? OFFSET ?`

	var response books.GetAllBookResponse
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	data := make([]books.Book, 0)
	for rows.Next() {
		var model books.BookModel
		err := rows.Scan(&model.ID, &model.Title, &model.Author, &model.Category, &model.PublicationYear, &model.TotalCopies, &model.AvailableCopies)
		if err != nil {
			return response, err
		}
		data = append(data, books.Book{
			ID:              model.ID,
			Title:           model.Title,
			Author:          model.Author,
			Category:        strings.Split(model.Category, ","),
			PublicationYear: model.PublicationYear,
			TotalCopies:     model.TotalCopies,
			AvailableCopies: model.AvailableCopies,
		})
	}
	response.Data = data
	response.Pagination = books.Pagination{
		Limit:  limit,
		Offset: offset,
	}
	return response, nil
}

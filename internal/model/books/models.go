package books

import "time"

type (
	AddBookRequest struct {
		Title           string   `json:"title"`
		Author          string   `json:"author"`
		Category        []string `json:"category"`
		PublicationYear string   `json:"publicationYear"`
		TotalCopies     int64    `json:"totalCopies"`
	}
)

type (
	BookModel struct {
		ID              int64     `db:"id"`
		Title           string    `db:"title"`
		Author          string    `db:"author"`
		Category        string    `db:"category"`
		PublicationYear string    `db:"publication_year"`
		TotalCopies     int64     `db:"total_copies"`
		AvailableCopies int64     `db:"available_copies"`
		CreatedAt       time.Time `db:"created_at"`
		UpdatedAt       time.Time `db:"updated_at"`
	}
)

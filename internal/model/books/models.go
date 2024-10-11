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

	DeleteBookRequest struct {
		Title  string `json:"title"`
		Author string `json:"author"`
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

type (
	GetAllBookResponse struct {
		Data       []Book     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Book struct {
		ID              int64    `json:"id"`
		Title           string   `json:"title"`
		Author          string   `json:"author"`
		Category        []string `json:"category"`
		PublicationYear string   `json:"publicationYear"`
		TotalCopies     int64    `json:"totalCopies"`
		AvailableCopies int64    `json:"availableCopies"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)

package memberships

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

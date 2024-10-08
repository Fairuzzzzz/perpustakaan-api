package memberships

import "time"

type SignUpRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserModel struct {
	ID        int64     `db:"id"`
	Email     string    `db:"email"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

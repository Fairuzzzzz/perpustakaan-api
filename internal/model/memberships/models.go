package memberships

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	DeleteUserRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
)

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
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

type (
	GetAllUserResponse struct {
		Data       []User     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	User struct {
		ID       int64  `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)

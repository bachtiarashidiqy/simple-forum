package memberships

import "time"

type (
	SingUpRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)

type (
	LoginResponse struct {
		AccessToken string `json:"access_token"`
	}
)

type (
	UserModel struct {
		ID        int64     `db:"id"`
		Username  string    `db:"username"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}
)

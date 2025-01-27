package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"post_title"`
		PostBody     string   `json:"post_body"`
		PostHashtags []string `json:"post_hashtags"`
	}
)

type (
	PostModel struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		PostTitle    string    `db:"post_title"`
		PostBody     string    `db:"post_body"`
		PostHashtags []string  `db:"post_hashtags"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
		CreatedBy    string    `db:"created_by"`
		UpdatedBy    string    `db:"updated_by"`
	}
)

package memberships

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bachtiarashidiqy/simple-forum/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `SELECT id, email, username, password, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? OR username = ?`
	row := r.db.QueryRowContext(ctx, query, email, username)
	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Username, &response.Password, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

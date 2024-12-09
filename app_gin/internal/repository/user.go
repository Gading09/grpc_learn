package repository

import (
	"context"
	"fmt"
	"gin/pkg/database"
	"gin/pkg/model"
	"gin/pkg/util"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	InsertUser(context.Context, model.User) error
	GetListUser(context.Context, util.Pagination, map[string]string) (result []model.User, count int64, err error)
	GetUserById(ctx context.Context, id string) (result model.User, err error)
	UpdateUser(ctx context.Context, user model.User) (err error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository() *userRepository {
	return &userRepository{
		db: database.DbPool,
	}
}

func (r *userRepository) InsertUser(ctx context.Context, user model.User) error {
	query := `
		INSERT INTO users (
			id,
			email,
			password,
			name,
			created_at,
			updated_at
		) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		ON CONFLICT (email)
		DO NOTHING
	`

	_, err := r.db.Exec(ctx, query, user.Id, user.Email, user.Password, user.Name, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetListUser(ctx context.Context, pagination util.Pagination, where map[string]string) (result []model.User, count int64, err error) {
	countQuery := "SELECT COUNT(*) FROM users WHERE deleted_at IS NULL"
	query := "SELECT id, email, name FROM users WHERE deleted_at IS NULL"

	for key, value := range where {
		if value != "" {
			nextQuery := fmt.Sprintf(` AND %s ILIKE '%%%s%%'`, key, value)
			countQuery += nextQuery
			query += nextQuery
		}
	}
	err = r.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		return
	}
	offset := (pagination.Page - 1) * pagination.Limit
	query += fmt.Sprintf(` LIMIT %d OFFSET %d`, pagination.Limit, offset)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Name); err != nil {
			return nil, 0, err
		}
		result = append(result, user)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return
}

func (r *userRepository) UpdateUser(ctx context.Context, user model.User) (err error) {
	query := `
		UPDATE 
			users 
		SET
			email = $1,
			password = $2,
			name = $3,
			updated_at = $4,
			deleted_at = $5
		WHERE 
			id = $6
	`

	_, err = r.db.Exec(ctx, query, user.Email, user.Password, user.Name, user.UpdatedAt, user.DeletedAt, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserById(ctx context.Context, id string) (result model.User, err error) {
	query := `
		SELECT
			id,
			email,
			password,
			name, 
			created_at,
			updated_at 
		FROM 
			users
		WHERE 
			id = $1 
			AND deleted_at IS NULL
	`

	err = r.db.QueryRow(ctx, query, id).Scan(&result.Id, &result.Email, &result.Password, &result.Name, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	return
}

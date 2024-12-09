package repository

import (
	"context"
	"fmt"
	"grpc/pkg/database"
	"grpc/pkg/model"
	"grpc/pkg/util"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	InsertUser(context.Context, model.User) error
	GetListUser(context.Context, util.Pagination, map[string]string) (result []model.User, count int64, err error)
	GetUserById(ctx context.Context, id string) (result model.User, err error)
	UpdateUser(ctx context.Context, user model.User) (err error)
}

//	type userRepository struct {
//		db *gorm.DB
//	}
type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository() *userRepository {
	return &userRepository{
		db: database.DbPool,
	}
}

// func (r userRepository) InsertUser(ctx context.Context, user model.User) (err error) {
// 	return r.db.Where(model.User{Email: user.Email}).Assign(model.User{
// 		Id:        user.Id,
// 		Email:     user.Email,
// 		Password:  user.Password,
// 		Name:      user.Name,
// 		CreatedAt: user.CreatedAt,
// 		UpdatedAt: user.UpdatedAt,
// 	}).FirstOrCreate(&model.User{}).Error
// }

// func (r userRepository) GetListUser(ctx context.Context, pagination util.Pagination, where map[string]string) (result []model.User, count int64, err error) {
// 	var field, sort string

// 	queryBuilder := r.db.Table("users").
// 		Select("id, email, name").
// 		Where("deleted_at isnull")

// 	if where["name"] != "" {
// 		name := fmt.Sprintf("%%%s%%", where["name"])
// 		queryBuilder = queryBuilder.Where("name ILIKE ?", name)
// 	}

// 	if where["email"] != "" {
// 		name := fmt.Sprintf("%%%s%%", where["email"])
// 		queryBuilder = queryBuilder.Where("email ILIKE ?", name)
// 	}

// 	if pagination.Field != "" {
// 		if pagination.Field == "name" {
// 			field = "name"
// 		} else if pagination.Field == "email" {
// 			field = "email"
// 		} else {
// 			field = "created_at"
// 		}
// 	} else {
// 		field = "created_at"
// 	}

// 	if pagination.Sort != "" {
// 		sort = pagination.Sort
// 	} else {
// 		sort = "DESC"
// 	}

// 	offset := (pagination.Page - 1) * pagination.Limit
// 	orderBy := fmt.Sprintf("%s %s", field, sort)
// 	limitBuilder := queryBuilder.Limit(int(pagination.Limit)).Offset(int(offset)).Order(orderBy)

// 	err = limitBuilder.Find(&result).Error
// 	if err != nil {
// 		return nil, count, err
// 	}

// 	err = queryBuilder.Count(&count).Error
// 	if err != nil {
// 		return nil, count, err
// 	}

// 	return
// }

// func (r userRepository) UpdateUser(ctx context.Context, user model.User) (err error) {
// 	return r.db.Save(&user).Error
// }

// func (r userRepository) GetUserById(ctx context.Context, id string) (result model.User, err error) {
// 	err = r.db.Where(model.User{Id: id}).First(&result).Error
// 	return
// }

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
			deleted_at = &5
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

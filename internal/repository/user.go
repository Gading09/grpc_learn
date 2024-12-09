package repository

import (
	"context"
	"fmt"
	"grpc/pkg/database"
	"grpc/pkg/model"
	"grpc/pkg/util"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	InsertUser(context.Context, model.User) error
	GetListUser(context.Context, util.Pagination, map[string]string) (result []model.User, count int64, err error)
	GetUserById(ctx context.Context, id string) (result model.User, err error)
	UpdateUser(ctx context.Context, user model.User) (err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() *userRepository {
	return &userRepository{
		db: database.DbGorm,
	}
}

func (r userRepository) InsertUser(ctx context.Context, user model.User) (err error) {
	return r.db.Where(model.User{Email: user.Email}).Assign(model.User{
		Id:        user.Id,
		Email:     user.Email,
		Password:  user.Password,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}).FirstOrCreate(&model.User{}).Error
}

func (r userRepository) GetListUser(ctx context.Context, pagination util.Pagination, where map[string]string) (result []model.User, count int64, err error) {
	var field, sort string

	queryBuilder := r.db.Table("users").
		Select("id, email, name").
		Where("deleted_at isnull")

	if where["name"] != "" {
		name := fmt.Sprintf("%%%s%%", where["name"])
		queryBuilder = queryBuilder.Where("name ILIKE ?", name)
	}

	if where["email"] != "" {
		name := fmt.Sprintf("%%%s%%", where["email"])
		queryBuilder = queryBuilder.Where("email ILIKE ?", name)
	}

	if pagination.Field != "" {
		if pagination.Field == "name" {
			field = "name"
		} else if pagination.Field == "email" {
			field = "email"
		} else {
			field = "created_at"
		}
	} else {
		field = "created_at"
	}

	if pagination.Sort != "" {
		sort = pagination.Sort
	} else {
		sort = "DESC"
	}

	offset := (pagination.Page - 1) * pagination.Limit
	orderBy := fmt.Sprintf("%s %s", field, sort)
	limitBuilder := queryBuilder.Limit(int(pagination.Limit)).Offset(int(offset)).Order(orderBy)

	err = limitBuilder.Find(&result).Error
	if err != nil {
		return nil, count, err
	}

	err = queryBuilder.Count(&count).Error
	if err != nil {
		return nil, count, err
	}

	return
}

func (r userRepository) UpdateUser(ctx context.Context, user model.User) (err error) {
	return r.db.Save(&user).Error
}

func (r userRepository) GetUserById(ctx context.Context, id string) (result model.User, err error) {
	err = r.db.Where(model.User{Id: id}).First(&result).Error
	return
}

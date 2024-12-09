package model

import "time"

type User struct {
	Id        string     `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type CreateUserResponse struct {
	Id string `json:"id"`
}

type GetListUserResponse struct {
	Page      int        `json:"page"`
	Limit     int        `json:"limit"`
	TotalPage int32      `json:"totalPage"`
	Total     int32      `json:"total"`
	ListUser  []ListUser `json:"listUser"`
}

type ListUser struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UpdateUserRequest struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type UpdateUserResponse struct {
	Id string `json:"id"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}

type DeleteUserResponse struct {
	Id string `json:"id"`
}

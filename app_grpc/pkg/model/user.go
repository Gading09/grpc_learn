package model

import "time"

type User struct {
	Id        string     `gorm:"primary_key" json:"id"`
	Email     string     `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password  string     `gorm:"size:255;not null" json:"password"`
	Name      string     `gorm:"size:255;not null" json:"name"`
	CreatedAt time.Time  `gorm:"" json:"created_at"`
	UpdatedAt time.Time  `gorm:"" json:"updated_at"`
	DeletedAt *time.Time `gorm:"" json:"deleted_at"`
}

package database

import (
	"context"
	"fmt"
	"grpc/pkg/model"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DbGorm *gorm.DB

func InitGorm(ctx context.Context) (*gorm.DB, error) {
	fmt.Println("===== Init DB =====")
	connectionString := os.Getenv("GORM_CONNECTION")
	db, err := gorm.Open("postgres", connectionString)
	db.LogMode(true)
	if err != nil || db.Error != nil {
		return nil, err
	}

	sqlDB := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil || db.Error != nil {
		return nil, err
	}

	DbGorm = db
	RunMigrations(DbGorm)
	return DbGorm, nil
}

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}

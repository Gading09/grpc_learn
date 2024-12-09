package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DbPool *pgxpool.Pool

func InitDB(ctx context.Context) (*pgxpool.Pool, error) {
	fmt.Println("===== Init DB =====")

	connectionString := os.Getenv("POSTGRE_CONNECTION")

	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %v", err)
	}

	config.MaxConns = 100
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 10 * time.Minute

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	DbPool = pool
	RunMigrate(ctx)

	return DbPool, nil
}

func RunMigrate(ctx context.Context) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		deleted_at TIMESTAMPTZ
	);
	`
	_, err := DbPool.Exec(ctx, createTableQuery)
	if err != nil {
		log.Fatalf("Error running migration: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}

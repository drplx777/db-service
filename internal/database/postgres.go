package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	cfg, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing database URL: %v", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}
	fmt.Println("Connected to the database successfully")

	CreateTable(pool) // Создание таблицы при подключении
	log.Println("Table creation check completed")
	return pool, nil
}

func CloseDB(pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
		fmt.Println("Database connection closed")
	} else {
		fmt.Println("No database connection to close")
	}

}

// CreateTable создаёт таблицу tasks с UUID PK и timestamp полями
func CreateTable(pool *pgxpool.Pool) error {
	query := `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    done_at TIMESTAMPTZ NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);	`

	if _, err := pool.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("ошибка при создании таблицы: %w", err)
	}
	log.Println("Таблица tasks успешно создана или уже существует")
	return nil
}

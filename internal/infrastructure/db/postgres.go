package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"go_web_api/internal/config"
)

func NewPostgresConnection(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	log.Println("PostgreSQL connection established")

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return db, nil
}

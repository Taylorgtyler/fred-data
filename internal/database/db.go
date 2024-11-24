package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/marcboeker/go-duckdb"
)

type DBContext struct {
	DB *sql.DB
}

func NewDBContext() (*DBContext, error) {
	token := os.Getenv("MOTHERDUCK_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("MOTHERDUCK_TOKEN environment variable is not set")
	}

	datasource := fmt.Sprintf("md:?motherduck_token=%s", token)
	db, err := sql.Open("duckdb", datasource)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MotherDuck: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping MotherDuck database: %w", err)
	}

	return &DBContext{DB: db}, nil
}

func (ctx *DBContext) Close() error {
	return ctx.DB.Close()
}

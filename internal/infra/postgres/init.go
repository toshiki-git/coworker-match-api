package postgres

import (
	"database/sql"
	"fmt"

	"github.com/coworker-match-api/internal/config"
	_ "github.com/lib/pq"
)

func Connect(config *config.LocalDBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}

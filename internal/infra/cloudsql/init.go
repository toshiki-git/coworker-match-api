package cloudsql

import (
	"database/sql"
	"fmt"

	"github.com/coworker-match-api/internal/config"
	_ "github.com/lib/pq"
)

// connectUnixSocket initializes a Unix socket connection pool for
// a Cloud SQL instance of PostgreSQL.
func ConnectUnixSocket(config *config.DevDBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		config.User, config.Password, config.DBName, config.UnixSocketPath)

	dbPool, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	err = dbPool.Ping()
	if err != nil {
		return nil, fmt.Errorf("dbPool.Ping: %w", err)
	}

	return dbPool, nil
}

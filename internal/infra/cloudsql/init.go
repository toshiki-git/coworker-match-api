package cloudsql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// connectUnixSocket initializes a Unix socket connection pool for
// a Cloud SQL instance of PostgreSQL.
func ConnectUnixSocket() (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_unix.go: %s environment variable not set.", k)
		}
		return v
	}

	var (
		dbUser         = mustGetenv("DB_USER")
		dbPwd          = mustGetenv("DB_PASSWORD")
		dbName         = mustGetenv("DB_NAME")
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET")
	)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		dbUser, dbPwd, dbName, unixSocketPath)

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

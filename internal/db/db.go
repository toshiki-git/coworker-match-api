package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	host := "postgres"     //os.Getenv("DB_HOST")
	user := "user"         //os.Getenv("DB_USER")
	password := "password" //os.Getenv("DB_PASSWORD")
	dbname := "db"         //os.Getenv("DB_NAME")
	sslmode := "disable"   //os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbname, sslmode)
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

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/coworker-match-api/internal/config"
	"github.com/coworker-match-api/internal/infra/cloudsql"
	"github.com/coworker-match-api/internal/infra/postgres"
	"github.com/coworker-match-api/internal/router"
)

func main() {
	// 設定をロード
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// データベース接続を初期化
	db, err := initializeDatabase(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// サーバー設定を初期化
	server := initializeServer(cfg, db)

	// サーバーを起動
	log.Printf("Server listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// データベース接続を初期化する関数
func initializeDatabase(cfg config.Config) (*sql.DB, error) {
	switch c := cfg.(type) {
	case *config.LocalConfig:
		return postgres.Connect(&c.DBConfig)
	case *config.DevConfig:
		return cloudsql.ConnectUnixSocket(&c.DBConfig)
	default:
		return nil, fmt.Errorf("unsupported environment: %v", c.GetEnvironment())
	}
}

// サーバー設定を初期化する関数
func initializeServer(cfg config.Config, db *sql.DB) *http.Server {
	var endPoint string
	var readTimeout time.Duration
	var writeTimeout time.Duration

	switch c := cfg.(type) {
	case *config.LocalConfig:
		endPoint = fmt.Sprintf(":%s", c.ServerConfig.Port)
		readTimeout = c.ServerConfig.ReadTimeout
		writeTimeout = c.ServerConfig.WriteTimeout
	case *config.DevConfig:
		endPoint = fmt.Sprintf(":%s", c.ServerConfig.Port)
		readTimeout = c.ServerConfig.ReadTimeout
		writeTimeout = c.ServerConfig.WriteTimeout
	}

	return &http.Server{
		Addr:         endPoint,
		Handler:      router.InitRouter(db),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
}

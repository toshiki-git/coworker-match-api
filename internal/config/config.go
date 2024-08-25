package config

import (
	"fmt"
	"log"
	"os"
)

type Config interface {
	GetEnvironment() string
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set. Program will exit.\n", key)
	}
	return value
}

// 環境変数を読み取り、Config構造体に設定する関数
func Load() (Config, error) {
	env := getEnv("ENV")

	switch env {
	case "local":
		return LoadLocalEnv()
	case "dev":
		return LoadDevEnv()
	default:
		return nil, fmt.Errorf("unknown environment: %s", env)
	}
}

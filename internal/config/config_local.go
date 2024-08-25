package config

import (
	"fmt"
	"strconv"
	"time"
)

type LocalDBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
}

type LocalServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type LocalRouterConfig struct {
	GoogleClientID string
}

type LocalConfig struct {
	Environment  string // local, dev, prod
	DBConfig     LocalDBConfig
	ServerConfig LocalServerConfig
	RouterConfig LocalRouterConfig
}

func LoadLocalEnv() (*LocalConfig, error) {
	cfg := &LocalConfig{}
	dbCfg := &LocalDBConfig{}
	serverCfg := &LocalServerConfig{}
	routerCfg := &LocalRouterConfig{}

	// Environment
	env := getEnv("ENV")
	cfg.Environment = env

	// DB Config
	dbCfg.Host = getEnv("DB_HOST")
	dbCfg.User = getEnv("DB_USER")
	dbCfg.Password = getEnv("DB_PASSWORD")
	dbCfg.DBName = getEnv("DB_NAME")

	// Server config
	serverCfg.Port = getEnv("API_PORT")

	readTimeout, err := strconv.Atoi(getEnv("READ_TIMEOUT"))
	if err != nil {
		return cfg, fmt.Errorf("invalid READ_TIMEOUT value")
	}
	serverCfg.ReadTimeout = time.Duration(readTimeout) * time.Second

	writeTimeout, err := strconv.Atoi(getEnv("WRITE_TIMEOUT"))
	if err != nil {
		return cfg, fmt.Errorf("invalid WRITE_TIMEOUT value")
	}
	serverCfg.WriteTimeout = time.Duration(writeTimeout) * time.Second

	// routerCfg
	routerCfg.GoogleClientID = getEnv("GOOGLE_CLIENT_ID")

	cfg.DBConfig = *dbCfg
	cfg.ServerConfig = *serverCfg
	cfg.RouterConfig = *routerCfg

	return cfg, nil
}

func (c *LocalConfig) GetEnvironment() string {
	return c.Environment
}

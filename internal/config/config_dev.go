package config

import (
	"fmt"
	"strconv"
	"time"
)

type DevDBConfig struct {
	UnixSocketPath string
	User           string
	Password       string
	DBName         string
}

type DevServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DevRouterConfig struct {
	GoogleClientID string
}

type DevConfig struct {
	Environment  string // local, dev, prod
	DBConfig     DevDBConfig
	ServerConfig DevServerConfig
	RouterConfig DevRouterConfig
}

func LoadDevEnv() (*DevConfig, error) {
	cfg := &DevConfig{}
	dbCfg := &DevDBConfig{}
	serverCfg := &DevServerConfig{}
	routerCfg := &DevRouterConfig{}

	// Environment
	env := getEnv("ENV")
	cfg.Environment = env

	// DB Config
	dbCfg.UnixSocketPath = getEnv("INSTANCE_UNIX_SOCKET")
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

func (c *DevConfig) GetEnvironment() string {
	return c.Environment
}

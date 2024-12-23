package config

import (
	"github.com/safayildirim/wallet-management-service/pkg/env"
)

type Config struct {
	App      AppConfig
	Http     HttpConfig
	Postgres PostgresConfig
}

var BaseConfig *Config

type AppConfig struct {
	AppEnv  string
	AppName string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	DBName          string
	MaxConn         int
	MaxIdleConn     int
	MaxLifeTimeConn int
	SslMode         string
}

type HttpConfig struct {
	Port int
	Host string
}

func init() {
	BaseConfig = New()
}

func New() *Config {
	return &Config{
		App: AppConfig{
			AppEnv:  env.New("APP_ENV", "dev").AsString(),
			AppName: env.New("APP_NAME", "wallet-management-service").AsString(),
		},
		Http: HttpConfig{
			Port: env.New("HTTP_PORT", "8080").AsInt(),
			Host: env.New("HTTP_HOST", "localhost").AsString(),
		},
		Postgres: PostgresConfig{
			Host:            env.New("PG_HOST", nil).AsString(),
			Port:            env.New("PG_PORT", nil).AsString(),
			User:            env.New("PG_USER", nil).AsString(),
			Pass:            env.New("PG_PASS", nil).AsString(),
			DBName:          env.New("PG_DB", nil).AsString(),
			MaxConn:         env.New("PG_MAX_CONNECTIONS", "10").AsInt(),
			MaxIdleConn:     env.New("PG_MAX_IDLE_CONNECTIONS", "10").AsInt(),
			MaxLifeTimeConn: env.New("PG_MAX_LIFETIME_CONNECTIONS", "20").AsInt(),
			SslMode:         env.New("PG_SSL_MODE", true).AsString(),
		},
	}
}

func IsDevEnv() bool {
	return BaseConfig.App.AppEnv == "dev"
}

func IsProdEnv() bool {
	return BaseConfig.App.AppEnv == "production"
}

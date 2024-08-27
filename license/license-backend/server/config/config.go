package config

import (
	"server/delivery/httpserver"
	"server/logger"
	"server/pkg/jwt"
	"server/repository/postgresql"
	"server/repository/sqlite"
	"server/service/log_service"
)

type MainConfig struct {
	PostgresqlCfg postgresql.Config  `yaml:"postgresql" env:"POSTGRESQL" `
	SqliteCfg     sqlite.Config      `yaml:"sqlite" env:"SQLITE" `
	JwtCfg        jwt.Config         `yaml:"jwt" env:"JWT"`
	LogCfg        logger.Config      `yaml:"log" env:"LOG"`
	ServerCfg     httpserver.Config  `yaml:"server" env:"SERVER"`
	ActivityLog   log_service.Config `yaml:"activity_log" env:"ACTIVITY_LOG"`
	DBType        string             `yaml:"db_type" env:"DB_TYPE" default:"sqlite"`
}

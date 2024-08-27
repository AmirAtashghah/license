package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var cfg MainConfig

func Load() *MainConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// db type

	cfg.DBType = os.Getenv("DB_TYPE")

	// postgresql

	cfg.PostgresqlCfg.Username = os.Getenv("POSTGRESQL_USERNAME")
	cfg.PostgresqlCfg.Password = os.Getenv("POSTGRESQL_PASSWORD")

	cfg.PostgresqlCfg.Port, _ = strconv.Atoi(os.Getenv("POSTGRESQL_PORT"))
	cfg.PostgresqlCfg.Host = os.Getenv("POSTGRESQL_HOST")
	cfg.PostgresqlCfg.DBName = os.Getenv("POSTGRESQL_DB_NAME")

	// sqlite

	cfg.SqliteCfg.DBName = os.Getenv("SQLITE_DB_NAME")

	// JWT Config

	cfg.JwtCfg.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

	// Logger Config

	cfg.LogCfg.FilePath = os.Getenv("LOG_FILE_PATH")
	cfg.LogCfg.UseLocalTime, _ = strconv.ParseBool(os.Getenv("LOG_USE_LOCAL_TIME"))
	cfg.LogCfg.FileMaxSizeInMB, _ = strconv.Atoi(os.Getenv("LOG_FILE_MAX_SIZE_IN_MB"))
	cfg.LogCfg.FileMaxAgeInDays, _ = strconv.Atoi(os.Getenv("LOG_FILE_MAX_AGE_IN_DAYS"))

	// Server Config

	cfg.ServerCfg.Host = os.Getenv("SERVER_HOST")
	cfg.ServerCfg.Port, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	cfg.ServerCfg.AllowPanel, _ = strconv.ParseBool(os.Getenv("ALLOW_PANEL"))

	// Activity log config
	cfg.ActivityLog.LogLanguage = os.Getenv("LOG_LANGUAGE")

	return &cfg
}

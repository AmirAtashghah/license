package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"server/logger"
)

const group = "postgresql"

//var Conn *pgx.Conn

type Config struct { // todo change tags
	Username string `yaml:"username" env:"USERNAME" `
	Password string `yaml:"password" env:"PASSWORD" `
	Port     int    `yaml:"port" env:"PORT" `
	Host     string `yaml:"host" env:"HOST" `
	DBName   string `yaml:"db_name" env:"DB_NAME" `
}

type PostgreSQLDB struct {
	config Config
	db     *pgxpool.Pool
}

func (p *PostgreSQLDB) Conn() *pgxpool.Pool {
	return p.db
}

func PostgreSQLConnection(cfg Config) *PostgreSQLDB {

	//log.Println(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
	//	cfg.Username,
	//	cfg.Password,
	//	cfg.Host,
	//	cfg.Port,
	//	cfg.DBName,
	//))

	db, err := pgxpool.New(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DBName,
		))

	if db == nil || err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error())

		log.Panic(err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	return &PostgreSQLDB{config: cfg, db: db}
}

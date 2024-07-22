package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

var Conn *pgx.Conn

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	DBName   string `koanf:"db_name"`
}

type PostgreSQLDB struct {
	config Config
	db     *pgx.Conn
}

func (p *PostgreSQLDB) Conn() *pgx.Conn {
	return p.db
}

func PostgreSQLConnection(cfg Config) *PostgreSQLDB {

	db, err := pgx.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DBName,
		))
	if err != nil {
		log.Panic(err)
	}
	return &PostgreSQLDB{config: cfg, db: db}
}

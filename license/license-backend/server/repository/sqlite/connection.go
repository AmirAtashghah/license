package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"server/pkg/hash"
	"strings"
	"time"
)

const group = "postgresql"

//var Conn *pgx.Conn

type Config struct { // todo change tags
	DBName string `yaml:"db_name" env:"DB_NAME" `
}

type SQLiteDB struct {
	config Config
	db     *sql.DB
}

func (p *SQLiteDB) Conn() *sql.DB {
	return p.db
}

func SQLiteConnection(cfg Config) *SQLiteDB {

	databasePath := fmt.Sprintf("/db/%s.sqlite", cfg.DBName) // deploy change /db/... push time

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	migrationFile := "./001_create_tables.sql" // deploy chang path ./001..
	migrationSQL, err := os.ReadFile(migrationFile)
	if err != nil {
		log.Fatalf("Failed to read migration file: %s\nError: %v", migrationFile, err)
	}

	// Split the migration SQL into individual statements
	queries := strings.Split(string(migrationSQL), ";")

	// Execute each query
	for _, query := range queries {
		// Trim any whitespace
		query = strings.TrimSpace(query)
		if query == "" {
			continue // Skip empty statements
		}
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Failed to execute query: %s\nError: %v", query, err)
		}
	}

	hashedPassword, err := hash.Hash("123")
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	insertUserQuery := `
	INSERT INTO users (name, username, password, role, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
	ON CONFLICT(username) DO NOTHING;
	`

	//_, err = db.Exec("DELETE FROM users")
	//if err != nil {
	//	log.Fatalf("failed to delete logs_template records: %w", err)
	//}

	_, err = db.Exec(insertUserQuery, "Super Admin", "root", hashedPassword, "superAdmin", time.Now().Unix(), -1)
	if err != nil {
		log.Fatalf("Failed to insert super user: %v", err)
	}

	_, err = db.Exec("DELETE FROM logs_template")
	if err != nil {
		log.Fatalf("failed to delete logs_template records: %w", err)
	}
	//_, err = db.Exec("DELETE FROM logs")
	//if err != nil {
	//	log.Fatalf("failed to delete logs records: %w", err)
	//}

	seedFile := "./seed.sql" // deploy change path ./s..
	seedSQL, err := os.ReadFile(seedFile)
	if err != nil {
		log.Fatalf("failed to read seed file: %w", err)
	}

	q := strings.Split(string(seedSQL), ";")
	for _, query := range q {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("failed to execute seed query: %w", err)
		}
	}

	return &SQLiteDB{config: cfg, db: db}
}

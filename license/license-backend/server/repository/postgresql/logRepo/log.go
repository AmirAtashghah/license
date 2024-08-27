package logRepo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"server/entity"
	"server/repository/postgresql"
)

var (
	insert = `
		INSERT INTO logs (title, message, created_at)
		VALUES ($1, $2, $3)
	`

	getByTitle = `
		SELECT id, title, message, created_at
		FROM logs
		WHERE title = $1
	`

	gets = `
		SELECT id, title, message, created_at
		FROM logs
	`

	insertTemplate = `
		INSERT INTO logs_template (key, value, language)
		VALUES ($1, $2, $3)
	`

	getTemplate = `
		SELECT id, key, value, language
		FROM logs_template
		WHERE key = $1 AND language = $2
	`

	getLicenseCheckByTitle = `
		SELECT id, title, message, created_at
		FROM logs
		WHERE title = $1
		AND message LIKE $2
		AND message LIKE $3;
	`
)

type DB struct {
	conn *postgresql.PostgreSQLDB
}

func New(conn *postgresql.PostgreSQLDB) *DB {
	return &DB{
		conn: conn,
	}
}

func (d *DB) Insert(logEntry *entity.Log) error {

	_, err := d.conn.Conn().Exec(context.Background(), insert,
		logEntry.Title, logEntry.Message, logEntry.CreatedAt,
	)

	return err
}

func (d *DB) GetByTitle(title string) ([]*entity.Log, error) {

	rows, err := d.conn.Conn().Query(context.Background(), getByTitle, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*entity.Log
	for rows.Next() {
		logEntry := new(entity.Log)
		err := rows.Scan(
			&logEntry.ID, &logEntry.Title, &logEntry.Message, &logEntry.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, logEntry)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return logs, nil
}

func (d *DB) Gets(offset, limit int16) ([]*entity.Log, error) {

	rows, err := d.conn.Conn().Query(context.Background(), gets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*entity.Log
	for rows.Next() {
		logEntry := new(entity.Log)
		err := rows.Scan(
			&logEntry.ID, &logEntry.Title, &logEntry.Message, &logEntry.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, logEntry)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return logs, nil
}

func (d *DB) GetLogTemplateByKeyAndLanguage(key, language string) (*entity.LogTemplate, error) {

	temp := new(entity.LogTemplate)

	log.Println(key, language)

	if err := d.conn.Conn().QueryRow(context.Background(),
		getTemplate, key, language).Scan(&temp.ID, &temp.Key,
		&temp.Value, &temp.Language); err != nil {

		log.Println(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return temp, nil
}

func (d *DB) GetLicenseCheckByTitle(customerID, productID string) ([]*entity.Log, error) {

	rows, err := d.conn.Conn().Query(context.Background(), getLicenseCheckByTitle, customerID, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*entity.Log
	for rows.Next() {
		logEntry := new(entity.Log)
		err := rows.Scan(
			&logEntry.ID, &logEntry.Title, &logEntry.Message, &logEntry.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, logEntry)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return logs, nil
}

func (d *DB) CountRecord() (int16, int16, int16) {

	var customers int16
	row := d.conn.Conn().QueryRow(context.Background(), "SELECT COUNT(*) FROM customers")
	if err := row.Scan(&customers); err != nil {
		return 0, 0, 0
	}
	var products int16
	rowP := d.conn.Conn().QueryRow(context.Background(), "SELECT COUNT(*) FROM products")
	if err := rowP.Scan(&products); err != nil {
		return 0, 0, 0
	}
	var users int16
	rowU := d.conn.Conn().QueryRow(context.Background(), "SELECT COUNT(*) FROM users")
	if err := rowU.Scan(&users); err != nil {
		return 0, 0, 0
	}

	return customers, products, users
}

package logRepo

import (
	"database/sql"
	"errors"
	"server/entity"
	"server/repository/sqlite"
)

var (
	insert = `
	INSERT INTO logs (title, message, created_at)
	VALUES (?, ?, ?)
`

	getByTitle = `
	SELECT id, title, message, created_at
	FROM logs
	WHERE title = ?
`

	gets = `
	SELECT id, title, message, created_at
	FROM logs
`

	insertTemplate = `
	INSERT INTO logs_template (key, value, language)
	VALUES (?, ?, ?)
`

	getTemplate = `
	SELECT id, key, value, language
	FROM logs_template
	WHERE key = ? AND language = ?
`

	getLicenseCheckByTitle = `
	SELECT id, title, message, created_at
	FROM logs
	WHERE title = ?
	AND message LIKE ?
	AND message LIKE ?;
`
)

type DB struct {
	conn *sqlite.SQLiteDB
}

func New(conn *sqlite.SQLiteDB) *DB {
	return &DB{
		conn: conn,
	}
}

func (d *DB) Insert(logEntry *entity.Log) error {

	_, err := d.conn.Conn().Exec(insert,
		logEntry.Title, logEntry.Message, logEntry.CreatedAt,
	)

	return err
}

func (d *DB) GetByTitle(title string) ([]*entity.Log, error) {

	rows, err := d.conn.Conn().Query(getByTitle, title)
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

	rows, err := d.conn.Conn().Query(gets)
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

	if err := d.conn.Conn().QueryRow(
		getTemplate, key, language).Scan(&temp.ID, &temp.Key,
		&temp.Value, &temp.Language); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return temp, nil
}

func (d *DB) GetLicenseCheckByTitle(customerID, productID string) ([]*entity.Log, error) {

	rows, err := d.conn.Conn().Query(getLicenseCheckByTitle, customerID, productID)
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
	row := d.conn.Conn().QueryRow("SELECT COUNT(*) FROM customers")
	if err := row.Scan(&customers); err != nil {
		return 0, 0, 0
	}
	var products int16
	rowP := d.conn.Conn().QueryRow("SELECT COUNT(*) FROM products")
	if err := rowP.Scan(&products); err != nil {
		return 0, 0, 0
	}
	var users int16
	rowU := d.conn.Conn().QueryRow("SELECT COUNT(*) FROM users")
	if err := rowU.Scan(&users); err != nil {
		return 0, 0, 0
	}

	return customers, products, users
}

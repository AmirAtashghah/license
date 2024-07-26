package postgresqllog

import (
	"context"
	"server/entity"
	"server/repository/postgresql"
)

const (
	INSERT = `
		INSERT INTO logs (id, client_hash, time, level, msg, "group", path, line, function, request_body, code)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	GET = `
		SELECT id, client_hash, time, level, msg, "group", path, line, function, request_body, code
		FROM logs
		WHERE client_hash = $1
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

func (d *DB) Insert(logEntry entity.Log) error {

	_, err := d.conn.Conn().Exec(context.Background(), INSERT,
		logEntry.ID, logEntry.ClientHash, logEntry.Time, logEntry.Level, logEntry.Message,
		logEntry.Group, logEntry.Path, logEntry.Line, logEntry.Function, logEntry.RequestBody, logEntry.Code,
	)

	return err
}

func (d *DB) GetByClientHash(clientHash string) ([]entity.Log, error) {

	rows, err := d.conn.Conn().Query(context.Background(), GET, clientHash)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []entity.Log
	for rows.Next() {
		var logEntry entity.Log
		err := rows.Scan(
			&logEntry.ID, &logEntry.ClientHash, &logEntry.Time, &logEntry.Level, &logEntry.Message,
			&logEntry.Group, &logEntry.Path, &logEntry.Line, &logEntry.Function, &logEntry.RequestBody, &logEntry.Code,
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

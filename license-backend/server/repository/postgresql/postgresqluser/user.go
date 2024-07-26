package postgresqluser

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"server/entity"
	"server/logger"
	"server/repository/postgresql"
)

const group = "postgresqluser"

var (
	GET = `SELECT id, username, password, login_at FROM users WHERE username=$1`
)

type DB struct {
	conn *postgresql.PostgreSQLDB
}

func New(conn *postgresql.PostgreSQLDB) *DB {
	return &DB{
		conn: conn,
	}
}

func (d *DB) GetByUsername(username string) (*entity.User, error) {
	if username == "" {
		logger.L().WithGroup(group).Error("error", "error", "username cannot be empty")

		return nil, fmt.Errorf("username cannot be empty")
	}

	row := d.conn.Conn().QueryRow(context.Background(), GET, username)

	var user entity.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.LoginAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		logger.L().WithGroup(group).Error("error", "error", err.Error())

		return nil, err
	}

	return &user, nil
}

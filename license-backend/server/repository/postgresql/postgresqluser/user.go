package postgresqluser

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"server/entity"
	"server/repository/postgresql"
)

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

func (d *DB) GetUserByUsername(username string) (*entity.User, error) {
	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}

	row := d.conn.Conn().QueryRow(context.Background(), GET, username)

	var user entity.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.LoginAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

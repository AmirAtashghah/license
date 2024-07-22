package postgresqluser

import "server/repository/postgresql"

type DB struct {
	conn *postgresql.PostgreSQLDB
}

func New(conn *postgresql.PostgreSQLDB) *DB {
	return &DB{
		conn: conn,
	}
}

func (d *DB) GetUserByUser() {}

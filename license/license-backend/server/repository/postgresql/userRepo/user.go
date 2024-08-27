package userRepo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"server/entity"
	"server/repository/postgresql"
)

var (
	insert  = `INSERT INTO users (name, username, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	update  = `UPDATE users SET name=$1, username=$2, password=$3, role=$4,created_at=$5 ,updated_at=$6 WHERE id=$7`
	get     = `SELECT id, name, username, password, role, created_at, updated_at FROM users WHERE username=$1`
	deleteU = `DELETE FROM users WHERE id=$1`
	gets    = `SELECT id, name, username, password, role, created_at, updated_at FROM users`
)

type DB struct {
	conn *postgresql.PostgreSQLDB
}

func New(conn *postgresql.PostgreSQLDB) *DB {
	return &DB{
		conn: conn,
	}
}

// todo add log

func (d *DB) Insert(user *entity.User) error {

	if _, err := d.conn.Conn().Exec(context.Background(),
		insert, user.Name, user.Username,
		user.Password, user.Role, user.CreatedAt,
		user.UpdatedAt); err != nil {

		log.Println(err)
		return err
	}

	return nil
}

func (d *DB) Update(user *entity.User) error {

	_, err := d.conn.Conn().Exec(context.Background(),
		update, user.Name, user.Username,
		user.Password, user.Role, user.CreatedAt, user.UpdatedAt,
		user.ID)
	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetByUsername(username string) (*entity.User, error) {

	user := new(entity.User)

	if err := d.conn.Conn().QueryRow(context.Background(),
		get, username).Scan(&user.ID, &user.Name,
		&user.Username, &user.Password, &user.Role,
		&user.CreatedAt, &user.UpdatedAt); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (d *DB) Delete(id int) error {

	_, err := d.conn.Conn().Exec(context.Background(), deleteU, id)
	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetAll() ([]*entity.User, error) {

	rows, err := d.conn.Conn().Query(context.Background(), gets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User

	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Password,
			&user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

package userRepo

import (
	"database/sql"
	"errors"
	"server/entity"
	"server/repository/sqlite"
)

var (

	// sqlite

	insert  = `INSERT INTO users (name, username, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	update  = `UPDATE users SET name=?, username=?, password=?, role=?, created_at=?, updated_at=? WHERE id=?`
	get     = `SELECT id, name, username, password, role, created_at, updated_at FROM users WHERE username=?`
	deleteL = `DELETE FROM users WHERE id=?`
	gets    = `SELECT id, name, username, password, role, created_at, updated_at FROM users`
	getByID = `SELECT id, name, username, password, role, created_at, updated_at FROM users WHERE id=?`
)

type DB struct {
	conn *sqlite.SQLiteDB
}

func New(conn *sqlite.SQLiteDB) *DB {
	return &DB{
		conn: conn,
	}
}

// todo add log

func (d *DB) Insert(user *entity.User) error {

	if _, err := d.conn.Conn().Exec(
		insert, user.Name, user.Username,
		user.Password, user.Role, user.CreatedAt,
		user.UpdatedAt); err != nil {

		return err
	}

	return nil
}

func (d *DB) Update(user *entity.User) error {

	_, err := d.conn.Conn().Exec(
		update, user.Name, user.Username,
		user.Password, user.Role, user.CreatedAt, user.UpdatedAt,
		user.ID)
	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetByID(id int) (*entity.User, error) {

	user := new(entity.User)

	if err := d.conn.Conn().QueryRow(
		getByID, id).Scan(&user.ID, &user.Name,
		&user.Username, &user.Password, &user.Role,
		&user.CreatedAt, &user.UpdatedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (d *DB) GetByUsername(username string) (*entity.User, error) {

	user := new(entity.User)

	if err := d.conn.Conn().QueryRow(
		get, username).Scan(&user.ID, &user.Name,
		&user.Username, &user.Password, &user.Role,
		&user.CreatedAt, &user.UpdatedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (d *DB) Delete(id int) error {

	_, err := d.conn.Conn().Exec(deleteL, id)
	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetAll() ([]*entity.User, error) {

	rows, err := d.conn.Conn().Query(gets)
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
			continue
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

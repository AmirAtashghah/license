package productRepo

import (
	"database/sql"
	"errors"
	"server/entity"
	"server/repository/sqlite"
)

var (

	// sqlite

	insert              = `INSERT INTO products (id, name, title, version, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	update              = `UPDATE products SET name=?, title=?, version=?, created_at=?, updated_at=? WHERE id=?`
	get                 = `SELECT id, name, title, version, created_at, updated_at FROM products WHERE id=?`
	deleteL             = `DELETE FROM products WHERE id=?`
	gets                = `SELECT id, name, title, version, created_at, updated_at FROM products`
	getByNameAndVersion = `SELECT id, name, title, version, created_at, updated_at FROM products WHERE name=? AND version=?`
)

type DB struct {
	conn *sqlite.SQLiteDB
}

func New(conn *sqlite.SQLiteDB) *DB {
	return &DB{
		conn: conn,
	}
}

//
//func (d *DB) GetByUsername(username string) (*entity.User, error) {
//	if username == "" {
//		logger.L().WithGroup(group).Error("error", "error", "username cannot be empty")
//
//		return nil, fmt.Errorf("username cannot be empty")
//	}
//
//	row := d.conn.Conn().QueryRow(    GET, username)
//
//	var user entity.User
//	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.LoginAt)
//	if err != nil {
//		if errors.Is(err,      sql.ErrNoRows) {
//			return nil, nil
//		}
//		logger.L().WithGroup(group).Error("error", "error", err.Error())
//
//		return nil, err
//	}
//
//	return &user, nil
//}

// todo add log

func (d *DB) Insert(product *entity.Product) error {

	if _, err := d.conn.Conn().Exec(insert,
		product.ID, product.Name,
		product.Title, product.Version,
		product.CreatedAt, product.UpdatedAt); err != nil {

		return err
	}

	return nil
}

func (d *DB) Update(product *entity.Product) error {

	_, err := d.conn.Conn().Exec(
		update, product.Name, product.Title,
		product.Version, product.CreatedAt,
		product.UpdatedAt, product.ID)

	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetByID(id string) (*entity.Product, error) {

	product := new(entity.Product)

	if err := d.conn.Conn().QueryRow(
		get, id).Scan(&product.ID, &product.Name,
		&product.Title, &product.Version,
		&product.CreatedAt, &product.UpdatedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return product, nil
}

func (d *DB) Delete(id string) error {

	_, err := d.conn.Conn().Exec(deleteL, id)
	if err != nil {

		return err
	}

	return nil
}

// todo add filters to query

func (d *DB) GetAll(limit, offset int16) ([]*entity.Product, error) {

	rows, err := d.conn.Conn().Query(gets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		product := new(entity.Product)
		err := rows.Scan(&product.ID, &product.Name, &product.Title, &product.Version, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			continue
		}
		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return products, nil
}

func (d *DB) GetProductByNameAndVersion(name string, version string) (*entity.Product, error) {

	product := new(entity.Product)
	if err := d.conn.Conn().QueryRow(
		getByNameAndVersion, name,
		version).Scan(&product.ID,
		&product.Name, &product.Title,
		&product.Version, &product.CreatedAt,
		&product.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return nil, nil
		}

		return nil, err
	}

	return product, nil
}

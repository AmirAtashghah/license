package productRepo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"server/entity"
	"server/repository/postgresql"
)

var (
	insert              = `INSERT INTO products (id, name, title, version, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	update              = `UPDATE products SET name=$1, title=$2, version=$3, created_at=$4, updated_at=$5 WHERE id=$6`
	get                 = `SELECT id, name, title, version, created_at, updated_at FROM products WHERE id=$1`
	deleteP             = `DELETE FROM products WHERE id=$1`
	gets                = `SELECT id, name, title, version, created_at, updated_at FROM products`
	getByNameAndVersion = `SELECT id, name, title, version, created_at, updated_at FROM products WHERE name=$1 AND version=$2`
)

type DB struct {
	conn *postgresql.PostgreSQLDB
}

func New(conn *postgresql.PostgreSQLDB) *DB {
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
//	row := d.conn.Conn().QueryRow(context.Background(), GET, username)
//
//	var user entity.User
//	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.LoginAt)
//	if err != nil {
//		if errors.Is(err, pgx.ErrNoRows) {
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

	if _, err := d.conn.Conn().Exec(context.Background(), insert,
		product.ID, product.Name,
		product.Title, product.Version,
		product.CreatedAt, product.UpdatedAt); err != nil {

		return err
	}

	return nil
}

func (d *DB) Update(product *entity.Product) error {

	_, err := d.conn.Conn().Exec(context.Background(),
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

	if err := d.conn.Conn().QueryRow(context.Background(),
		get, id).Scan(&product.ID, &product.Name,
		&product.Title, &product.Version,
		&product.CreatedAt, &product.UpdatedAt); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return product, nil
}

func (d *DB) Delete(id string) error {

	_, err := d.conn.Conn().Exec(context.Background(), deleteP, id)
	if err != nil {

		return err
	}

	return nil
}

// todo add filters to query

func (d *DB) GetAll(limit, offset int16) ([]*entity.Product, error) {

	rows, err := d.conn.Conn().Query(context.Background(), gets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		product := new(entity.Product)
		err := rows.Scan(&product.ID, &product.Name, &product.Title, &product.Version, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
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
	if err := d.conn.Conn().QueryRow(context.Background(),
		getByNameAndVersion, name,
		version).Scan(&product.ID,
		&product.Name, &product.Title,
		&product.Version, &product.CreatedAt,
		&product.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {

			return nil, nil
		}

		return nil, err
	}

	return product, nil
}

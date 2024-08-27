package customerRepo

import (
	"database/sql"
	"errors"
	"server/entity"
	"server/repository/sqlite"
)

var (

	// sqlite
	insert            = `INSERT INTO customers (id, name, email, phone, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	update            = `UPDATE customers SET name=?, email=?, phone=?, updated_at=?, created_at=? WHERE id=?`
	get               = `SELECT id, name, email, phone, created_at, updated_at FROM customers WHERE id=?`
	deleteL           = `DELETE FROM customers WHERE id=?`
	getByNameAndEmail = `SELECT id, name, email, phone, created_at, updated_at FROM customers WHERE name=? AND email=?`
	gets              = `SELECT id, name, email, phone, created_at, updated_at FROM customers`
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

func (d *DB) Insert(customer *entity.Customer) error {

	if _, err := d.conn.Conn().Exec(insert,
		customer.ID, customer.Name,
		customer.Email, customer.Phone,
		customer.CreatedAt, customer.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (d *DB) Update(customer *entity.Customer) error {

	_, err := d.conn.Conn().Exec(
		update, customer.Name,
		customer.Email, customer.Phone,
		customer.CreatedAt, customer.UpdatedAt,
		customer.ID)

	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetByID(id string) (*entity.Customer, error) {

	customer := new(entity.Customer)

	if err := d.conn.Conn().QueryRow(
		get, id).Scan(&customer.ID, &customer.Name,
		&customer.Email, &customer.Phone,
		&customer.CreatedAt, &customer.UpdatedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customer, nil
}

func (d *DB) Delete(id string) error {

	_, err := d.conn.Conn().Exec(deleteL, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetCustomerByNameAndEmail(name, email string) (*entity.Customer, error) {

	customer := new(entity.Customer)
	err := d.conn.Conn().QueryRow(getByNameAndEmail, name,
		email).Scan(&customer.ID, &customer.Name,
		&customer.Email, &customer.Phone,
		&customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customer, nil
}

func (d *DB) GetAll(limit, offset int16) ([]*entity.Customer, error) {

	rows, err := d.conn.Conn().Query(gets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*entity.Customer

	for rows.Next() {
		customer := new(entity.Customer)
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			continue
		}
		customers = append(customers, customer)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return customers, nil
}

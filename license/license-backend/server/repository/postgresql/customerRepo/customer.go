package customerRepo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"server/entity"
	"server/repository/postgresql"
)

var (
	// postgresql
	insert            = `INSERT INTO customers (id, name, email, phone, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	update            = `UPDATE customers SET name=$1, email=$2, phone=$3, updated_at=$4, created_at=$5 WHERE id=$6`
	get               = `SELECT id, name, email, phone, created_at, updated_at FROM customers WHERE id=$1`
	deleteC           = `DELETE FROM customers WHERE id=$1`
	getByNameAndEmail = `SELECT id, name, email, phone, created_at, updated_at FROM customers WHERE name=$1 AND email=$2`
	gets              = `SELECT id, name, email, phone, created_at, updated_at FROM customers`
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

func (d *DB) Insert(customer *entity.Customer) error {

	log.Println(customer)
	if _, err := d.conn.Conn().Exec(context.Background(), insert,
		customer.ID, customer.Name,
		customer.Email, customer.Phone,
		customer.CreatedAt, customer.UpdatedAt); err != nil {
		log.Println("40", err)
		return err
	}

	return nil
}

func (d *DB) Update(customer *entity.Customer) error {

	_, err := d.conn.Conn().Exec(context.Background(),
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

	if err := d.conn.Conn().QueryRow(context.Background(),
		get, id).Scan(&customer.ID, &customer.Name,
		&customer.Email, &customer.Phone,
		&customer.CreatedAt, &customer.UpdatedAt); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customer, nil
}

func (d *DB) Delete(id string) error {

	_, err := d.conn.Conn().Exec(context.Background(), deleteC, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (d *DB) GetCustomerByNameAndEmail(name, email string) (*entity.Customer, error) {

	customer := new(entity.Customer)
	err := d.conn.Conn().QueryRow(context.Background(), getByNameAndEmail, name,
		email).Scan(&customer.ID, &customer.Name,
		&customer.Email, &customer.Phone,
		&customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customer, nil
}

func (d *DB) GetAll(limit, offset int16) ([]*entity.Customer, error) {

	rows, err := d.conn.Conn().Query(context.Background(), gets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*entity.Customer

	for rows.Next() {
		customer := new(entity.Customer)
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		customers = append(customers, customer)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return customers, nil
}

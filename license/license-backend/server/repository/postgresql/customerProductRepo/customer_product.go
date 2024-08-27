package customerProductRepo

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
	insert                      = `INSERT INTO customer_products (id,customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	update                      = `UPDATE customer_products SET customer_id=$1, product_id=$2, hardware_hash=$3, license_type=$4, is_active=$5, expire_at=$6, first_confirmed_at=$7, last_confirmed_at=$8, created_at=$9, updated_at=$10 WHERE id=$11`
	get                         = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at, updated_at FROM customer_products WHERE id=$1`
	deleteP                     = `DELETE FROM customer_products WHERE id=$1`
	getByCustomerIDAndProductID = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at,updated_at FROM customer_products WHERE customer_id=$1 AND product_id=$2`
	gets                        = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at, updated_at FROM customer_products`
	getByCustomerID             = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at,updated_at FROM customer_products WHERE customer_id=$1`
	getByProductID              = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at,updated_at FROM customer_products WHERE product_id=$1`
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

func (d *DB) Insert(customerProduct *entity.CustomerProduct) error {

	if _, err := d.conn.Conn().Exec(context.Background(), insert, customerProduct.ID,
		customerProduct.CustomerID, customerProduct.ProductID,
		customerProduct.HardwareHash, customerProduct.LicenseType,
		customerProduct.IsActive, customerProduct.ExpireAt,
		customerProduct.FirstConfirmedAt, customerProduct.LastConfirmedAt,
		customerProduct.CreatedAt, customerProduct.UpdatedAt); err != nil {

		log.Println(err)
		return err
	}

	return nil
}

func (d *DB) Update(customerProduct *entity.CustomerProduct) error {

	_, err := d.conn.Conn().Exec(context.Background(),
		update, customerProduct.CustomerID,
		customerProduct.ProductID, customerProduct.HardwareHash,
		customerProduct.LicenseType, customerProduct.IsActive,
		customerProduct.ExpireAt, customerProduct.FirstConfirmedAt,
		customerProduct.LastConfirmedAt, customerProduct.CreatedAt, customerProduct.UpdatedAt,
		customerProduct.ID)

	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetByID(id string) (*entity.CustomerProduct, error) {

	customerProduct := new(entity.CustomerProduct)

	if err := d.conn.Conn().QueryRow(context.Background(),
		get, id).Scan(&customerProduct.ID, &customerProduct.CustomerID,
		&customerProduct.ProductID, &customerProduct.HardwareHash,
		&customerProduct.LicenseType, &customerProduct.IsActive,
		&customerProduct.ExpireAt, &customerProduct.FirstConfirmedAt,
		&customerProduct.LastConfirmedAt, &customerProduct.CreatedAt,
		&customerProduct.UpdatedAt); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customerProduct, nil
}

func (d *DB) Delete(id string) error {

	_, err := d.conn.Conn().Exec(context.Background(), deleteP, id)
	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetByCustomerIDAndProductID(customerID string, productID string) (*entity.CustomerProduct, error) {

	customerProduct := new(entity.CustomerProduct)
	err := d.conn.Conn().QueryRow(context.Background(), getByCustomerIDAndProductID, customerID, productID).Scan(
		&customerProduct.ID,
		&customerProduct.CustomerID,
		&customerProduct.ProductID,
		&customerProduct.HardwareHash,
		&customerProduct.LicenseType,
		&customerProduct.IsActive,
		&customerProduct.ExpireAt,
		&customerProduct.FirstConfirmedAt,
		&customerProduct.LastConfirmedAt,
		&customerProduct.CreatedAt,
		&customerProduct.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customerProduct, nil
}

func (d *DB) GetAll(limit, offset int16) ([]*entity.CustomerProduct, error) {

	rows, err := d.conn.Conn().Query(context.Background(), gets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customerProducts []*entity.CustomerProduct

	for rows.Next() {
		cp := new(entity.CustomerProduct)
		err := rows.Scan(&cp.ID, &cp.CustomerID, &cp.ProductID, &cp.HardwareHash,
			&cp.LicenseType, &cp.IsActive, &cp.ExpireAt,
			&cp.FirstConfirmedAt, &cp.LastConfirmedAt,
			&cp.CreatedAt, &cp.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		customerProducts = append(customerProducts, cp)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return customerProducts, nil

}

func (d *DB) GetByCustomerID(customerID string) (*entity.CustomerProduct, error) {

	customerProduct := new(entity.CustomerProduct)
	err := d.conn.Conn().QueryRow(context.Background(), getByCustomerID, customerID).Scan(
		&customerProduct.ID,
		&customerProduct.CustomerID,
		&customerProduct.ProductID,
		&customerProduct.HardwareHash,
		&customerProduct.LicenseType,
		&customerProduct.IsActive,
		&customerProduct.ExpireAt,
		&customerProduct.FirstConfirmedAt,
		&customerProduct.LastConfirmedAt,
		&customerProduct.CreatedAt,
		&customerProduct.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customerProduct, nil
}

func (d *DB) GetByProductID(productID string) (*entity.CustomerProduct, error) {

	customerProduct := new(entity.CustomerProduct)
	err := d.conn.Conn().QueryRow(context.Background(), getByProductID, productID).Scan(
		&customerProduct.ID,
		&customerProduct.CustomerID,
		&customerProduct.ProductID,
		&customerProduct.HardwareHash,
		&customerProduct.LicenseType,
		&customerProduct.IsActive,
		&customerProduct.ExpireAt,
		&customerProduct.FirstConfirmedAt,
		&customerProduct.LastConfirmedAt,
		&customerProduct.CreatedAt,
		&customerProduct.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customerProduct, nil
}

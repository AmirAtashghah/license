package customerProductRepo

import (
	"database/sql"
	"errors"
	"server/entity"
	"server/repository/sqlite"
)

var (
	insert                      = `INSERT INTO customer_products (id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	update                      = `UPDATE customer_products SET customer_id=?, product_id=?, hardware_hash=?, license_type=?, is_active=?, expire_at=?, first_confirmed_at=?, last_confirmed_at=?, created_at=?, updated_at=? WHERE id=?`
	get                         = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at, updated_at FROM customer_products WHERE id=?`
	deleteC                     = `DELETE FROM customer_products WHERE id=?`
	getByCustomerIDAndProductID = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at, updated_at FROM customer_products WHERE customer_id=? AND product_id=?`
	gets                        = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at, updated_at FROM customer_products`
	getByCustomerID             = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at, updated_at FROM customer_products WHERE customer_id=?`
	getByProductID              = `SELECT id, customer_id, product_id, hardware_hash, license_type, is_active, expire_at, first_confirmed_at, last_confirmed_at, created_at, updated_at FROM customer_products WHERE product_id=?`
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

func (d *DB) Insert(customerProduct *entity.CustomerProduct) error {

	if _, err := d.conn.Conn().Exec(insert, customerProduct.ID,
		customerProduct.CustomerID, customerProduct.ProductID,
		customerProduct.HardwareHash, customerProduct.LicenseType,
		customerProduct.IsActive, customerProduct.ExpireAt,
		customerProduct.FirstConfirmedAt, customerProduct.LastConfirmedAt,
		customerProduct.CreatedAt, customerProduct.UpdatedAt); err != nil {

		return err
	}

	return nil
}

func (d *DB) Update(customerProduct *entity.CustomerProduct) error {

	_, err := d.conn.Conn().Exec(
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

	if err := d.conn.Conn().QueryRow(
		get, id).Scan(&customerProduct.ID, &customerProduct.CustomerID,
		&customerProduct.ProductID, &customerProduct.HardwareHash,
		&customerProduct.LicenseType, &customerProduct.IsActive,
		&customerProduct.ExpireAt, &customerProduct.FirstConfirmedAt,
		&customerProduct.LastConfirmedAt, &customerProduct.CreatedAt,
		&customerProduct.UpdatedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customerProduct, nil
}

func (d *DB) Delete(id string) error {

	_, err := d.conn.Conn().Exec(deleteC, id)
	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetByCustomerIDAndProductID(customerID string, productID string) (*entity.CustomerProduct, error) {

	customerProduct := new(entity.CustomerProduct)
	err := d.conn.Conn().QueryRow(getByCustomerIDAndProductID, customerID, productID).Scan(
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customerProduct, nil
}

func (d *DB) GetAll(limit, offset int16) ([]*entity.CustomerProduct, error) {

	rows, err := d.conn.Conn().Query(gets)
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
	err := d.conn.Conn().QueryRow(getByCustomerID, customerID).Scan(
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customerProduct, nil
}

func (d *DB) GetByProductID(productID string) (*entity.CustomerProduct, error) {

	customerProduct := new(entity.CustomerProduct)
	err := d.conn.Conn().QueryRow(getByProductID, productID).Scan(
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return customerProduct, nil
}

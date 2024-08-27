package restrictionRepo

import (
	"database/sql"
	"errors"
	"server/entity"
	"server/pkg/param"
	"server/repository/sqlite"
)

// todo check all querys

var (

	// customer product restrictions
	insertCPR                                     = `INSERT INTO customer_products_restrictions (restriction_id, customer_product_id, value, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	updateCPR                                     = `UPDATE customer_products_restrictions SET restriction_id=?, customer_product_id=?, value=?, created_at=?,updated_at=? WHERE id=?`
	getCPR                                        = `SELECT id,restriction_id, customer_product_id, value, created_at, updated_at FROM customer_products_restrictions WHERE id=?`
	deleteCPR                                     = `DELETE FROM customer_products_restrictions WHERE customer_product_id=?`
	getByCustomerRestrictionIDAndRestrictionIDCPR = `SELECT id,restriction_id, customer_product_id, value, created_at, updated_at FROM customer_products_restrictions WHERE customer_product_id=? AND restriction_id=?`
	getAllCPR                                     = `SELECT id,restriction_id, customer_product_id, value, created_at, updated_at FROM customer_products_restrictions`
	getByCustomersProductIDCPR                    = `SELECT id,restriction_id, customer_product_id, value, created_at, updated_at FROM customer_products_restrictions WHERE customer_product_id=? `

	// restrictions

	insert   = `INSERT INTO restrictions (key, created_at, updated_at) VALUES (?, ?, ?)`
	update   = `UPDATE restrictions SET key=?, created_at=?, updated_at=? WHERE id=?`
	get      = `SELECT id,key, created_at, updated_at FROM restrictions WHERE id=?`
	deleteL  = `DELETE FROM restrictions WHERE id=?`
	getByKey = `SELECT id, key, created_at, updated_at FROM restrictions WHERE key=?`
	getAll   = `SELECT id, key, created_at, updated_at FROM restrictions`
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

func (d *DB) Insert(restriction *entity.Restriction) error {

	if _, err := d.conn.Conn().Exec(insert,
		restriction.Key,
		restriction.CreatedAt, restriction.UpdatedAt); err != nil {

		return err
	}

	return nil
}

func (d *DB) Update(restriction *entity.Restriction) error {

	_, err := d.conn.Conn().Exec(
		update,
		restriction.Key, restriction.CreatedAt,
		restriction.UpdatedAt, restriction.ID)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetByID(id string) (*entity.Restriction, error) {

	restriction := new(entity.Restriction)

	if err := d.conn.Conn().QueryRow(
		get, id).Scan(&restriction.ID,
		&restriction.Key, &restriction.CreatedAt,
		&restriction.UpdatedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return restriction, nil
}

func (d *DB) Delete(id int16) error {

	_, err := d.conn.Conn().Exec(deleteL, id)
	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetByKey(key string) (*entity.Restriction, error) {

	restriction := new(entity.Restriction)
	err := d.conn.Conn().QueryRow(getByKey,
		key).Scan(&restriction.ID, &restriction.Key,
		&restriction.CreatedAt, &restriction.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return restriction, nil
}

//
//func (d *DB) GetByProductID(productID string) ([]*entity.Restriction, error) {
//
//	rows, err := d.conn.Conn().Query(getByProductID, productID)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var restrictions []*entity.Restriction
//
//	for rows.Next() {
//		restriction := new(entity.Restriction)
//		err := rows.Scan(&restriction.ID, &restriction.ProductID, &restriction.Key, &restriction.Value, &restriction.CreatedAt, &restriction.UpdatedAt)
//		if err != nil {
//			return nil, err
//		}
//		restrictions = append(restrictions, restriction)
//	}
//
//	if rows.Err() != nil {
//		return nil, rows.Err()
//	}
//
//	return restrictions, nil
//}

func (d *DB) GetAll(filter *param.GetAllRestrictionsRequest) ([]*entity.Restriction, error) {

	rows, err := d.conn.Conn().Query(getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restrictions []*entity.Restriction

	for rows.Next() {
		restriction := new(entity.Restriction)
		err := rows.Scan(&restriction.ID, &restriction.Key, &restriction.CreatedAt, &restriction.UpdatedAt)
		if err != nil {
			return nil, err
		}
		restrictions = append(restrictions, restriction)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return restrictions, nil
}

///////// customer product restriction

func (d *DB) InsertCPR(cpr *entity.CustomersProductRestriction) error {

	//log.Println(cpr)

	if _, err := d.conn.Conn().Exec(insertCPR,
		cpr.RestrictionID, cpr.CustomersProductID, cpr.Value, cpr.CreatedAt, cpr.UpdatedAt); err != nil {

		return err
	}

	return nil
}

func (d *DB) UpdateCPR(cpr *entity.CustomersProductRestriction) error {

	_, err := d.conn.Conn().Exec(
		updateCPR,
		cpr.RestrictionID, cpr.CustomersProductID, cpr.Value, cpr.CreatedAt, cpr.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetByIDCPR(id string) (*entity.CustomersProductRestriction, error) {

	cpr := new(entity.CustomersProductRestriction)

	if err := d.conn.Conn().QueryRow(
		getCPR, id).Scan(&cpr.ID, &cpr.RestrictionID, &cpr.CustomersProductID, &cpr.Value, &cpr.CreatedAt, &cpr.UpdatedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return cpr, nil
}

func (d *DB) DeleteCPR(customerProductID string) error {

	_, err := d.conn.Conn().Exec(deleteCPR, customerProductID)
	if err != nil {

		return err
	}

	return nil
}

func (d *DB) GetByCustomersProductIDAndRestrictionIDCPR(customersProductID, restrictionID string) (*entity.CustomersProductRestriction, error) {

	cpr := new(entity.CustomersProductRestriction)
	err := d.conn.Conn().QueryRow(getByCustomerRestrictionIDAndRestrictionIDCPR,
		customersProductID, restrictionID).Scan(&cpr.ID, &cpr.RestrictionID,
		&cpr.CustomersProductID, &cpr.Value, &cpr.CreatedAt, &cpr.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return cpr, nil
}

func (d *DB) GetByCustomersProductIDCPR(customersProductID string) ([]*entity.CustomersProductRestriction, error) {

	rows, err := d.conn.Conn().Query(getByCustomersProductIDCPR, customersProductID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cprs []*entity.CustomersProductRestriction

	for rows.Next() {
		cpr := new(entity.CustomersProductRestriction)
		err := rows.Scan(&cpr.ID, &cpr.RestrictionID, &cpr.CustomersProductID, &cpr.Value, &cpr.CreatedAt, &cpr.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cprs = append(cprs, cpr)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return cprs, nil
}

func (d *DB) GetAllCPR(filter *param.GetAllRestrictionsRequest) ([]*entity.CustomersProductRestriction, error) {

	rows, err := d.conn.Conn().Query(getAllCPR)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cprs []*entity.CustomersProductRestriction

	for rows.Next() {
		cpr := new(entity.CustomersProductRestriction)
		err := rows.Scan(&cpr.ID, &cpr.RestrictionID, &cpr.CustomersProductID, &cpr.Value, &cpr.CreatedAt, &cpr.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cprs = append(cprs, cpr)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return cprs, nil
}

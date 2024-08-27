package restrictionRepo

import (
	"server/repository/postgresql"
)

// todo check all querys

var (
	insert               = `INSERT INTO restrictions (product_id, key, value, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	update               = `UPDATE restrictions SET product_id=$1, key=$2, value=$3,created_at=$4, updated_at=$5 WHERE id=$6`
	get                  = `SELECT id, product_id, key, value, created_at, updated_at FROM restrictions WHERE id=$1`
	deleteR              = `DELETE FROM restrictions WHERE id=$1`
	getByProductIDAndKey = `SELECT id, product_id, key, value, created_at, updated_at FROM restrictions WHERE product_id=$1 AND key=$2`
	getByProductID       = `SELECT id, product_id, key, value, created_at, updated_at FROM restrictions WHERE product_id=$1`
	getAll               = `SELECT id, product_id, key, value, created_at, updated_at FROM restrictions`
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
//
//func (d *DB) Insert(restriction *entity.Restriction) error {
//
//	if _, err := d.conn.Conn().Exec(context.Background(), insert,
//		restriction.ProductID, restriction.Key,
//		restriction.Value,
//		restriction.CreatedAt, restriction.UpdatedAt); err != nil {
//
//		return err
//	}
//
//	return nil
//}
//
//func (d *DB) Update(restriction *entity.Restriction) error {
//
//	log.Println(restriction)
//	_, err := d.conn.Conn().Exec(context.Background(),
//		update, restriction.ProductID,
//		restriction.Key, restriction.Value, restriction.CreatedAt,
//		restriction.UpdatedAt, restriction.ID)
//
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	return nil
//}
//
//func (d *DB) GetByID(id string) (*entity.Restriction, error) {
//
//	restriction := new(entity.Restriction)
//
//	if err := d.conn.Conn().QueryRow(context.Background(),
//		get, id).Scan(&restriction.ID, &restriction.ProductID,
//		&restriction.Key, &restriction.Value, &restriction.CreatedAt,
//		&restriction.UpdatedAt); err != nil {
//
//		if errors.Is(err, pgx.ErrNoRows) {
//			return nil, nil
//		}
//
//		return nil, err
//	}
//
//	return restriction, nil
//}
//
//func (d *DB) Delete(id int16) error {
//
//	_, err := d.conn.Conn().Exec(context.Background(), deleteR, id)
//	if err != nil {
//
//		return err
//	}
//
//	return nil
//}
//
//func (d *DB) GetByProductIDAndKey(productID, key string) (*entity.Restriction, error) {
//
//	restriction := new(entity.Restriction)
//	err := d.conn.Conn().QueryRow(context.Background(), getByProductIDAndKey,
//		productID, key).Scan(&restriction.ID,
//		&restriction.ProductID, &restriction.Key,
//		&restriction.Value,
//		&restriction.CreatedAt, &restriction.UpdatedAt)
//	if err != nil {
//		if errors.Is(err, pgx.ErrNoRows) {
//			return nil, nil
//		}
//
//		return nil, err
//	}
//
//	return restriction, nil
//}
//
//func (d *DB) GetByProductID(productID string) ([]*entity.Restriction, error) {
//
//	rows, err := d.conn.Conn().Query(context.Background(), getByProductID, productID)
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
//
//func (d *DB) GetAll(filter *param.GetAllRestrictionsRequest) ([]*entity.Restriction, error) {
//
//	rows, err := d.conn.Conn().Query(context.Background(), getAll)
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

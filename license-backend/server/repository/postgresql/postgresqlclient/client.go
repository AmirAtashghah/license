package postgresqlclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"server/entity"
	"server/pkg/param"
	"server/repository/postgresql"
	"strings"
)

type DB struct {
	conn *postgresql.PostgreSQLDB
}

func New(conn *postgresql.PostgreSQLDB) *DB {
	return &DB{
		conn: conn,
	}
}

const (
	INSERT = `INSERT INTO clients (id, software_name, software_version, hardware_hash, license_type, user_metadata, is_active, expires_at, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	GET    = `SELECT id, software_name, software_version, hardware_hash, license_type, user_metadata, is_active, expires_at, created_at, updated_at, deleted_at FROM clients WHERE id = $1`
	GETS   = `SELECT id, software_name, software_version, hardware_hash, license_type, user_metadata, is_active, expires_at, created_at, updated_at, deleted_at FROM clients WHERE 1=1`
	DELETE = `DELETE FROM clients WHERE id=$1`
	UPDATE = `UPDATE clients SET `
)

func (d *DB) Insert(client entity.Client) error {

	_, err := d.conn.Conn().Exec(context.Background(), INSERT, client.ID, client.SoftwareName, client.SoftwareVersion, client.HardwareHash, client.LicenseType, client.UserMetadata, client.IsActive, client.ExpiresAt, client.CreatedAt, client.UpdatedAt, client.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetByID(id string) (*entity.Client, error) {

	var client entity.Client

	err := d.conn.Conn().QueryRow(context.Background(), GET, id).Scan(
		&client.ID, &client.SoftwareName, &client.SoftwareVersion, &client.HardwareHash,
		&client.LicenseType, &client.UserMetadata, &client.ExpiresAt, &client.CreatedAt,
		&client.UpdatedAt, &client.DeletedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &client, nil
}

func (d *DB) GetAll(filter param.ClientFilter) ([]entity.Client, error) {

	var queryBuilder strings.Builder
	queryBuilder.WriteString(GETS)

	var args []interface{}
	argID := 1

	if filter.LicenseType != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND license_type=$%d", argID))
		args = append(args, *filter.LicenseType)
		argID++
	}
	if filter.UserMetadata != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND user_metadata=$%d", argID))
		args = append(args, *filter.UserMetadata)
		argID++
	}

	if filter.IsActive != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND is_active=$%d", argID))
		args = append(args, *filter.IsActive)
		argID++
	}

	if filter.Limit != nil {
		queryBuilder.WriteString(fmt.Sprintf(" LIMIT $%d", argID))
		args = append(args, *filter.Limit)
		argID++
	}

	if filter.Offset != nil {
		queryBuilder.WriteString(fmt.Sprintf(" OFFSET $%d", argID))
		args = append(args, *filter.Offset)
		argID++
	}

	query := queryBuilder.String()

	rows, err := d.conn.Conn().Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []entity.Client
	for rows.Next() {
		var client entity.Client
		err := rows.Scan(
			&client.ID,
			&client.SoftwareName,
			&client.SoftwareVersion,
			&client.HardwareHash,
			&client.LicenseType,
			&client.UserMetadata,
			&client.ExpiresAt,
			&client.CreatedAt,
			&client.UpdatedAt,
			&client.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return clients, nil
}

func (d *DB) Update(id string, client entity.Client) error {
	if id == "" {
		return fmt.Errorf("client ID cannot be empty")
	}

	var queryBuilder strings.Builder
	queryBuilder.WriteString(UPDATE)
	var args []interface{}
	argID := 1

	if client.SoftwareName != nil {
		queryBuilder.WriteString(fmt.Sprintf("software_name=$%d,", argID))
		args = append(args, *client.SoftwareName)
		argID++
	}
	if client.SoftwareVersion != nil {
		queryBuilder.WriteString(fmt.Sprintf("software_version=$%d,", argID))
		args = append(args, *client.SoftwareVersion)
		argID++
	}
	if client.HardwareHash != nil {
		queryBuilder.WriteString(fmt.Sprintf("hardware_hash=$%d,", argID))
		args = append(args, *client.HardwareHash)
		argID++
	}
	if client.LicenseType != nil {
		queryBuilder.WriteString(fmt.Sprintf("license_type=$%d,", argID))
		args = append(args, *client.LicenseType)
		argID++
	}
	if client.UserMetadata != nil {
		queryBuilder.WriteString(fmt.Sprintf("user_metadata=$%d,", argID))
		args = append(args, *client.UserMetadata)
		argID++
	}
	if client.IsActive != nil {
		queryBuilder.WriteString(fmt.Sprintf("is_active=$%d,", argID))
		args = append(args, *client.IsActive)
		argID++
	}
	if client.ExpiresAt != nil {
		queryBuilder.WriteString(fmt.Sprintf("expires_at=$%d,", argID))
		args = append(args, *client.ExpiresAt)
		argID++
	}
	if client.UpdatedAt != nil {
		queryBuilder.WriteString(fmt.Sprintf("updated_at=$%d,", argID))
		args = append(args, *client.UpdatedAt)
		argID++
	}

	// Remove the trailing comma
	query := strings.TrimSuffix(queryBuilder.String(), ",")

	// Append the WHERE clause
	queryBuilder.Reset()
	queryBuilder.WriteString(query)
	queryBuilder.WriteString(fmt.Sprintf(" WHERE id=$%d", argID))
	args = append(args, id)

	// Execute the query
	_, err := d.conn.Conn().Exec(context.Background(), queryBuilder.String(), args...)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) Delete(id string) error {
	if id == "" {
		return fmt.Errorf("client ID cannot be empty")
	}

	_, err := d.conn.Conn().Exec(context.Background(), DELETE, id)
	if err != nil {
		return err
	}

	return nil
}

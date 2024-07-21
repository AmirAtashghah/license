package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"server/entity"
)

const (
	EXISTS = `SELECT EXISTS(SELECT 1 FROM clients WHERE id=$1)`
	INSERT = `
		INSERT INTO clients (
			id, software_name, software_version, hardware_hash, license_type, user_metadata, expires_at, created_at, updated_at, deleted_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		)
	`
	GET = `
		SELECT id, software_name, software_version, hardware_hash, license_type, user_metadata, expires_at, created_at, updated_at, deleted_at
		FROM clients
		WHERE id = $1
	`
)

func ClientExists(conn *pgx.Conn, id string) (bool, error) {

	var exists bool

	err := conn.QueryRow(context.Background(), EXISTS, id).Scan(&exists)

	return exists, err
}

func InsertClient(conn *pgx.Conn, client entity.Client) error {

	_, err := conn.Exec(context.Background(), INSERT, client.ID, client.SoftwareName, client.SoftwareVersion, client.HardwareHash, client.LicenseType, client.UserMetadata, client.ExpiresAt, client.CreatedAt, client.UpdatedAt, client.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func GetClientByID(conn *pgx.Conn, id string) (*entity.Client, error) {

	var client entity.Client

	err := conn.QueryRow(context.Background(), GET, id).Scan(
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

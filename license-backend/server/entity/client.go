package entity

type Client struct {
	ID              string  `json:"id"`
	SoftwareName    *string `json:"software_name"`
	SoftwareVersion *string `json:"software_version"`
	HardwareHash    *string `json:"hardware_hash"`
	LicenseType     *string `json:"license_type"`
	UserMetadata    *string `json:"user_metadata"`
	IsActive        *bool   `json:"is_active"`
	ExpiresAt       *int64  `json:"expires_at"`
	CreatedAt       *int64  `json:"created_at"`
	UpdatedAt       *int64  `json:"updated_at"`
	DeletedAt       *int64  `json:"deleted_at"`
}

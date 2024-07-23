package param

type ClientRequest struct {
	ID              string `json:"id"`
	SoftwareName    string `json:"software_name" validate:"required"`
	SoftwareVersion string `json:"software_version" validate:"required"`
	HardwareHash    string `json:"hardware_hash" validate:"required"`
	LicenseType     string `json:"license_type" validate:"required"`
	UserMetadata    string `json:"user_metadata" validate:"required"`
	TimeStamp       int64  `json:"time_stamp" validate:"required,numeric"`
	RandomNumber    int64  `json:"random_number" validate:"required,numeric"`
}

type ClientResponse struct {
	ClientID         string `json:"client_id"`
	AuthKey          string `json:"auth_key"`
	ValidationStatus bool   `json:"validation_status"`
}

type UpdateClientRequest struct {
	ID           string `json:"id"`
	LicenseType  string `json:"license_type"`
	UserMetadata string `json:"user_metadata"`
	ExpiresAt    int64  `json:"expires_at"`
}

type DeleteClientRequest struct {
	ID string `json:"id"`
}

type ChangeActivateRequest struct {
	ID         string `json:"id"`
	IsActivate bool   `json:"is_activate"`
}

type ClientFilter struct {
	Limit        *int    `json:"limit"`
	Offset       *int    `json:"offset"`
	LicenseType  *string `json:"license_type"`
	UserMetadata *string `json:"user_metadata"`
	IsActive     *bool   `json:"is_active"`
}

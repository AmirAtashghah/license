package param

type ClientRequest struct {
	ID              string `json:"id"`
	SoftwareName    string `json:"software_name" validate:"required"`
	SoftwareVersion string `json:"software_version" validate:"required"`
	HardwareHash    string `json:"hardware_hash" validate:"required,sha256"`
	LicenseType     string `json:"license_type" validate:"required"`
	UserMetadata    string `json:"user_metadata" validate:"required"`
	TimeStamp       int64  `json:"time_stamp" validate:"required,numeric"`
	RandomNumber    int64  `json:"random_number" validate:"required,numeric"`
}

type ClientResponse struct {
	AuthKey          int64 `json:"auth_key"`
	ValidationStatus bool  `json:"validation_status"`
}

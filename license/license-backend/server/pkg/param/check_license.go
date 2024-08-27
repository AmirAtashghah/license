package param

type CheckLicenseRequest struct {
	ID           string `json:"id" validate:"required"`
	HardwareHash string `json:"hardwareHash" validate:"required"`
	//ServerTimestamp int64  `json:"serverTimestamp" validate:"required,numeric"`
	TimeStamp    int64 `json:"timeStamp" validate:"required,numeric"`
	RandomNumber int64 `json:"randomNumber" validate:"required,numeric"`
}

type CheckLicenseResponse struct {
	AuthKey          string `json:"auth_key"`
	Restriction      string `json:"restriction"`
	ValidationStatus bool   `json:"validationStatus"`
}

type EncryptedCheckLicenseRequest struct {
	Body string `json:"body"`
}

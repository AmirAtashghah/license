package entity

type CustomerProduct struct {
	ID               string `json:"id"`
	CustomerID       string `json:"customerID"`
	ProductID        string `json:"productID"`
	HardwareHash     string `json:"hardwareHash"`
	LicenseType      string `json:"licenseType"`
	IsActive         bool   `json:"isActive"`
	ExpireAt         int64  `json:"expireAt"`
	FirstConfirmedAt int64  `json:"firstConfirmedAt"`
	LastConfirmedAt  int64  `json:"lastConfirmedAt"`
	CreatedAt        int64  `json:"createdAt"`
	UpdatedAt        int64  `json:"updatedAt"`
}

package param

type CreateCustomerProductRequest struct {
	CustomerID   string `json:"customerID" validate:"required"`
	ProductID    string `json:"productID" validate:"required"`
	Restrictions string `json:"restrictions"` // string  json [{restrictionID:value},{restrictionID1:value1}]
	LicenseType  string `json:"licenseType" validate:"required"`
	IsActive     *bool  `json:"isActive" validate:"required"`
	ExpireAt     int64  `json:"expireAt" validate:"required"`
}

type UpdateCustomerProductRequest struct {
	ID               string `json:"id" validate:"required"`
	CustomerID       string `json:"customerID" validate:"required"`
	ProductID        string `json:"productID" validate:"required"`
	Restrictions     string `json:"restrictions"` // string  json [{restrictionID:value},{restrictionID1:value1}]
	HardwareHash     string `json:"hardwareHash" validate:"required"`
	LicenseType      string `json:"licenseType" validate:"required"`
	IsActive         *bool  `json:"isActive" validate:"required"`
	ExpireAt         int64  `json:"expireAt" validate:"required"`
	FirstConfirmedAt int64  `json:"firstConfirmedAt" validate:"required"`
	LastConfirmedAt  int64  `json:"lastConfirmedAt" validate:"required"`
	CreatedAt        int64  `json:"createdAt" validate:"required"`
}

type DeleteCustomerProductRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetCustomerProductRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetAllCustomerProductsRequest struct {
	Limit  int16 `json:"limit"`
	Offset int16 `json:"offset"`
}

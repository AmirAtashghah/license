package param

type CreateRestrictionRequest struct {
	Key string `json:"key" validate:"required"`
}

type UpdateRestrictionRequest struct {
	ID        int16  `json:"id" validate:"required"`
	Key       string `json:"key" validate:"required"`
	CreatedAt int64  `json:"createdAt" validate:"required"`
}

type DeleteRestrictionRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetRestrictionRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetAllRestrictionsRequest struct {
	Limit  int16 `json:"limit"`
	Offset int16 `json:"offset"`
}

type GetRestrictionsByProductRequest struct { /// todo remove this
	ProductID string `json:"product_id"`
}

// //

type CreateCustomersProductRestrictionRequest struct {
	RestrictionIDAndValues string `json:"restriction_id_and_values" validate:"required"`
	CustomersProductID     string `json:"customers_product_id" validate:"required"`
}

type DeleteCustomersProductRestrictionRequest struct {
	CustomersProductID string `json:"customers_product_id" validate:"required"`
}

type GetCustomersProductRestrictionRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetAllCustomersProductRestrictionsRequest struct {
	Limit  int16 `json:"limit"`
	Offset int16 `json:"offset"`
}

type GetCustomersProductRestrictionsByCPAndRestrictionIDRequest struct {
	RestrictionID      string `json:"restriction_id" validate:"required"`
	CustomersProductID string `json:"customers_product_id" validate:"required"`
}

type GetCustomersProductRestrictionsByCustomerProductIDRequest struct {
	CustomersProductID string `json:"customers_product_id" validate:"required"`
}

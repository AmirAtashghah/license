package param

type CreateCustomerRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

type UpdateCustomerRequest struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	CreatedAt int64  `json:"createdAt" validate:"required"`
}

type DeleteCustomerRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetCustomerRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetAllCustomersRequest struct {
	Limit  int16 `json:"limit"`
	Offset int16 `json:"offset"`
}

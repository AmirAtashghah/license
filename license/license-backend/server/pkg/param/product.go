package param

type CreateProductRequest struct {
	Name    string `json:"name" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Version string `json:"version" validate:"required"`
}

type UpdateProductRequest struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Version   string `json:"version" validate:"required"`
	CreatedAt int64  `json:"createdAt" validate:"required"`
}

type DeleteProductRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetProductRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetAllProductsRequest struct {
	Limit  int16 `json:"limit"`
	Offset int16 `json:"offset"`
}

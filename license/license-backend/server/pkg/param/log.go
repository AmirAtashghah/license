package param

type GetActivityLogsByTitleRequest struct {
	Title    string `json:"title" validate:"required"`
	Language string `json:"language" validate:"required"`
}

type GetLicenseCheckActivityLogsByTitleRequest struct {
	ProductID  string `json:"productID" validate:"required"`
	CustomerID string `json:"customerID" validate:"required"`
	Language   string `json:"language" validate:"required"`
}

type GetActivityLogsRequest struct {
	Limit    int16  `json:"limit"`
	Offset   int16  `json:"offset"`
	Language string `json:"language"`
}

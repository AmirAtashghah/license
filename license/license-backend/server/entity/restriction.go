package entity

type Restriction struct {
	ID        int16  `json:"id"`
	Key       string `json:"key"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type CustomersProductRestriction struct {
	ID                 int16  `json:"id"`
	RestrictionID      int16  `json:"restrictionID"`
	CustomersProductID string `json:"customersProductID"`
	Value              string `json:"value"`
	CreatedAt          int64  `json:"createdAt"`
	UpdatedAt          int64  `json:"updatedAt"`
}

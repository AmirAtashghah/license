package param

type ValidationErr struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

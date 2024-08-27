package entity

type Log struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"createdAt"`
}

type LogTemplate struct {
	ID       int64  `json:"id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Language string `json:"language"`
}

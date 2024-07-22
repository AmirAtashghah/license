package entity

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	LoginAt  int64  `json:"login_at"`
}

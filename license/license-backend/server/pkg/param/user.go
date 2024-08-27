package param

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ValidateTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UpdateUserRequest struct {
	ID        int    `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" `
	Role      string `json:"role" validate:"required"`
	CreatedAt int64  `json:"createdAt" validate:"required"`
}

type DeleteUserRequest struct {
	ID int `json:"id" validate:"required"`
}

type GetUserRequest struct {
	Username string `json:"username" validate:"required"`
}

package dto

type ErrorResponse struct {
	Message string `json:"message"`
}

type OKResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

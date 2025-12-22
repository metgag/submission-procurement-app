package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"omitempty,oneof=admin user"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateSupplierRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address"`
}

type UpdateSupplierRequest struct {
	Name    *string `json:"name,omitempty" validate:"omitempty,min=2"`
	Email   *string `json:"email,omitempty" validate:"omitempty,email"`
	Address *string `json:"address,omitempty"`
}

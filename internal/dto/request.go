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

type CreateItemRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type CreatePurchasingItemRequest struct {
	SupplierItemID uint `json:"supplier_item_id" validate:"required"`
	Quantity       int  `json:"quantity" validate:"required,gt=0"`
}

type CreateSupplierItemRequest struct {
	SupplierID uint  `json:"supplier_id" validate:"required"`
	ItemID     uint  `json:"item_id" validate:"required"`
	Price      int64 `json:"price" validate:"required,gt=0"`
	Stock      int   `json:"stock" validate:"required,gte=0"`
}

type UpdateSupplierItemRequest struct {
	Price int64 `json:"price" validate:"omitempty,gt=0"`
	Stock int   `json:"stock" validate:"omitempty,gte=0"`
}

type CreatePurchasingRequest struct {
	SupplierID uint                          `json:"supplier_id" validate:"required"`
	Items      []CreatePurchasingItemRequest `json:"items" validate:"required,min=1,dive"`
}

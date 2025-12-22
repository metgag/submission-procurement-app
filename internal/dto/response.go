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

type SupplierResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type ItemResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type SupplierItemResponse struct {
	ID           uint   `json:"id"`
	SupplierID   uint   `json:"supplier_id"`
	ItemID       uint   `json:"item_id"`
	Price        int64  `json:"price"`
	Stock        int    `json:"stock"`
	ItemName     string `json:"item_name,omitempty"`
	SupplierName string `json:"supplier_name,omitempty"`
}

type PurchasingItemResponse struct {
	ItemID   uint  `json:"item_id"`
	Quantity int   `json:"quantity"`
	Subtotal int64 `json:"subtotal"`
}

type PurchasingResponse struct {
	ID         uint                     `json:"id"`
	GrandTotal int64                    `json:"grand_total"`
	Items      []PurchasingItemResponse `json:"items"`
}

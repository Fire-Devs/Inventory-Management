package models

type Inventory struct {
	ID          int
	Name        string     `json:"name" validate:"required"`
	Stock       int        `json:"stock" validate:"required"`
	Description string     `json:"description"`
	CoverImage  string     `json:"cover_image"`
	Category    []Category `json:"category"`
	Supplier    []Supplier `json:"supplier"`
	Prices      []Prices   `json:"prices"`
}

type Prices struct {
	ID          int
	Currency    string `json:"currency" validate:"required"`
	InventoryID int    `json:"inventory_id"`
	Amount      int    `json:"amount" validate:"required"`
}
type Category struct {
	ID          int
	Name        string `json:"name" validate:"required"`
	InventoryID int    `json:"inventory_id"`
	Description string `json:"description" validate:"required"`
}

type Supplier struct {
	ID          int
	Name        string `json:"name" validate:"required"`
	InventoryID int    `json:"inventory_id"`
	ContactInfo string `json:"contact_info" validate:"required"`
}

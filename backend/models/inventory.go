package models

type Inventory struct {
	ID         int
	Name       string `json:"name" validate:"required"`
	Price      int    `json:"price" validate:"required"`
	Stock      int    `json:"stock" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
	CoverImage string `json:"cover_image"`
	SupplierID int    `json:"supplier_id"`
}

type InventoryData struct {
	ID          int
	Description string                 `bson:"description" json:"description"`
	MetaData    map[string]interface{} `bson:"meta_data" json:"meta_data"`
	Images      []map[string]string    `bson:"images" json:"images"`
	Size        []map[string]string    `bson:"size" json:"size"`
	Colors      []string               `bson:"colors" json:"colors"`
}

type Category struct {
	ID          int
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type Supplier struct {
	ID          int
	Name        string `json:"name" validate:"required"`
	ContactInfo string `json:"contact_info" validate:"required"`
}

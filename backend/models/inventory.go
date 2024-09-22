package models

import "go.mongodb.org/mongo-driver/bson"

type Inventory struct {
	ID            int
	Name          string        `json:"name" validate:"required"`
	Price         int           `json:"price" validate:"required"`
	Stock         int           `json:"stock" validate:"required"`
	CategoryID    int           `json:"category_id" validate:"required"`
	CoverImage    string        `json:"cover_image"`
	SupplierID    int           `json:"supplier_id"`
	InventoryData InventoryData `json:"inventory_data"`
}

type InventoryData struct {
	ID          int
	Description string   `bson:"description" json:"description"`
	MetaData    bson.M   `bson:"meta_data" json:"meta_data"`
	Images      []string `bson:"images" json:"images"`
	Features    []string `bson:"features" json:"features"`
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

package repository

import (
	"InventoryManagement/database"
	"InventoryManagement/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertCategory(category *models.Category) error {
	postgres, _ := database.Connect()

	_, err := postgres.Exec(context.Background(),
		"INSERT INTO categories (name, description) VALUES ($1, $2)",
		category.Name, category.Description)
	if err != nil {
		return err
	}

	return nil
}

func GetCategories() ([]models.Category, error) {
	postgres, _ := database.Connect()

	rows, err := postgres.Query(context.Background(), "SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
func InsertSupplier(supplier *models.Supplier) error {
	postgres, _ := database.Connect()

	_, err := postgres.Exec(context.Background(),
		"INSERT INTO suppliers (name, contact_info) VALUES ($1, $2)",
		supplier.Name, supplier.ContactInfo)
	if err != nil {
		return err
	}

	return nil
}

func InsertInventory(inventory *models.Inventory) error {
	postgres, _ := database.Connect()
	mongo, _ := database.ConnectMongo()

	id, err := postgres.Exec(context.Background(),
		"INSERT INTO inventory (name, stock, price, cover_image, category_id, supplier_id) VALUES ($1, $2) RETURNING id",
		inventory.Name, inventory.Stock, inventory.Price, inventory.CoverImage, inventory.CategoryID, inventory.SupplierID)
	if err != nil {
		return err
	}

	_, err = mongo.Database("inventory").Collection("inventory_data").InsertOne(context.Background(),
		bson.M{
			"inventory_id": id,
			"description":  inventory.InventoryData.Description,
			"meta_data":    inventory.InventoryData.MetaData,
			"images":       inventory.InventoryData.Images,
			"features":     inventory.InventoryData.Features,
		})

	if err != nil {
		return err
	}

	return nil

}

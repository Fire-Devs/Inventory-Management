package repository

import (
	"InventoryManagement/database"
	"InventoryManagement/models"
	"context"
	"fmt"
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

func GetSuppliers() ([]models.Supplier, error) {
	postgres, _ := database.Connect()

	rows, err := postgres.Query(context.Background(), "SELECT * FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []models.Supplier
	for rows.Next() {
		var supplier models.Supplier
		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.ContactInfo)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func InsertInventory(inventory *models.Inventory) error {
	postgres, _ := database.Connect()
	mongo, _ := database.ConnectMongo()

	var id int
	err := postgres.QueryRow(context.Background(),
		"INSERT INTO inventory (name, stock, price, cover_image, category_id, supplier_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		inventory.Name, inventory.Stock, inventory.Price, inventory.CoverImage, inventory.CategoryID, inventory.SupplierID).Scan(&id)
	if err != nil {
		return err
	}

	fmt.Println(id)

	_, err = mongo.Database("inventory").Collection("inventory_data").InsertOne(context.Background(),
		bson.D{
			{Key: "inventory_id", Value: id},
			{Key: "description", Value: inventory.InventoryData.Description},
			{Key: "meta_data", Value: inventory.InventoryData.MetaData},
			{Key: "images", Value: inventory.InventoryData.Images},
			{Key: "features", Value: inventory.InventoryData.Features},
		})

	if err != nil {
		return err
	}
	return nil
}

func GetInventory() ([]models.Inventory, error) {
	postgres, _ := database.Connect()

	rows, err := postgres.Query(context.Background(), "SELECT id, name, stock, price, cover_image, category_id, supplier_id FROM inventory")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []models.Inventory
	for rows.Next() {
		var inventory models.Inventory
		err := rows.Scan(&inventory.ID, &inventory.Name, &inventory.Stock, &inventory.Price, &inventory.CoverImage, &inventory.CategoryID, &inventory.SupplierID)
		if err != nil {
			return nil, err
		}

		mongo, _ := database.ConnectMongo()
		err = mongo.Database("inventory").Collection("inventory_data").FindOne(context.Background(), bson.M{"inventory_id": inventory.ID}).Decode(&inventory.InventoryData)
		if err != nil {
			return nil, err
		}

		inventories = append(inventories, inventory)
	}

	return inventories, nil
}

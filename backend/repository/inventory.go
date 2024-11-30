package repository

import (
	"InventoryManagement/database"
	"InventoryManagement/models"
	"context"
	"fmt"
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

	var id int
	err := postgres.QueryRow(context.Background(),
		"INSERT INTO inventory (name, stock, cover_image) VALUES ($1, $2, $3) RETURNING id",
		inventory.Name, inventory.Stock, inventory.CoverImage).Scan(&id)
	if err != nil {

		return err
	}

	if len(inventory.Prices) != 0 {
		for _, price := range inventory.Prices {
			_, err = postgres.Exec(context.Background(),
				"INSERT INTO price (currency, amount, inventory) VALUES ($1, $2, $3)",
				price.Currency, price.Amount, id)
			if err != nil {
				fmt.Println("Price Error")
				return err
			}
		}
	}

	if len(inventory.Category) != 0 {
		for _, category := range inventory.Category {
			_, err = postgres.Exec(context.Background(),
				"INSERT INTO categories (name, description, inventory) VALUES ($1, $2, $3)",
				category.Name, category.Description, id)
			if err != nil {
				fmt.Println("Category Error")
				return err
			}
		}
	}

	if len(inventory.Supplier) != 0 {
		for _, supplier := range inventory.Supplier {
			_, err = postgres.Exec(context.Background(),
				"INSERT INTO suppliers (name, contact_info, inventory) VALUES ($1, $2, $3)",
				supplier.Name, supplier.ContactInfo, id)
			if err != nil {
				fmt.Println("Supplier Error")
				return err
			}
		}
	}

	return nil
}

func GetInventory() ([]models.Inventory, error) {
	postgres, _ := database.Connect()

	rows, err := postgres.Query(context.Background(), "SELECT id, name, stock, cover_image FROM inventory")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []models.Inventory
	for rows.Next() {
		var inventory models.Inventory
		err := rows.Scan(&inventory.ID, &inventory.Name, &inventory.Stock, &inventory.CoverImage)
		if err != nil {
			fmt.Println("Inventory Error" + err.Error())
			return nil, err

		}
		inventories = append(inventories, inventory)
	}

	for i, inventory := range inventories {

		rows, err := postgres.Query(context.Background(), "SELECT id, currency, amount, inventory FROM price WHERE inventory = $1", inventory.ID)
		if err != nil {
			return nil, err

		}

		var prices []models.Prices
		for rows.Next() {
			var price models.Prices
			err := rows.Scan(&price.ID, &price.Currency, &price.Amount, &price.InventoryID)
			if err != nil {
				fmt.Println("Price Error")
				return nil, err
			}
			prices = append(prices, price)
		}
		inventories[i].Prices = prices

		rows, err = postgres.Query(context.Background(), "SELECT id, name, description, inventory FROM categories WHERE inventory = $1", inventory.ID)
		if err != nil {
			return nil, err
		}

		var categories []models.Category
		for rows.Next() {
			var category models.Category
			err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.InventoryID)
			if err != nil {
				fmt.Println("Category Error")
				return nil, err
			}
			categories = append(categories, category)
		}
		inventories[i].Category = categories

		rows, err = postgres.Query(context.Background(), "SELECT id, name, contact_info, inventory FROM suppliers WHERE inventory = $1", inventory.ID)
		if err != nil {
			return nil, err
		}

		var suppliers []models.Supplier
		for rows.Next() {
			var supplier models.Supplier
			err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.ContactInfo, &supplier.InventoryID)
			if err != nil {
				fmt.Println("Supplier Error")
				return nil, err
			}
			suppliers = append(suppliers, supplier)
		}
		inventories[i].Supplier = suppliers

	}

	defer rows.Close()

	return inventories, nil

}

func InsertPrice(prices *models.Prices) (id string, err error) {
	postgres, _ := database.Connect()
	err = postgres.QueryRow(context.Background(),
		"INSERT INTO price (currency, amount, inventory) VALUES ($1, $2, $3) returning id",
		prices.Currency, prices.Amount).Scan()
	if err != nil {
		return "", err
	}
	return id, nil
}

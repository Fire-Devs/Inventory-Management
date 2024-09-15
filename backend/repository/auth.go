package repository

import (
	"InventoryManagement/database"
	"context"
)

func CreateUser(name string, email string, password string) error {
	conn, err := database.Connect()
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", name, email, password)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func GetUserByEmailOrName(data string) (string, string, string, error) {
	conn, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var name, email, password string
	err = conn.QueryRow(context.Background(), "SELECT name, email, password FROM users WHERE name = $1 OR email = $1", data).Scan(&name, &email, &password)
	if err != nil {
		panic(err)
		return "", "", "", err
	}

	return name, email, password, nil
}

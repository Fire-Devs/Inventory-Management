package repository

import (
	"InventoryManagement/database"
	"InventoryManagement/models"
	"context"
	"time"
)

func CreateUser(name string, email string, password string) error {
	conn, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var id int
	err = conn.QueryRow(context.Background(), "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", name, email, password).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmailOrName(user *models.UserLogin) (string, string, string, error) {
	conn, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var data string
	if user.Email != "" {
		data = user.Email
	} else {
		data = user.Name
	}

	var name, email, password string
	err = conn.QueryRow(context.Background(), "SELECT name, email, password FROM users WHERE name = $1 OR email = $1", data).Scan(&name, &email, &password)
	if err != nil {
		return "", "", "", err
	}

	return name, email, password, nil
}

func SetUserToken(id string, token string) error {
	rdb := database.ConnectRedis()
	ctx := context.Background()

	_, err := rdb.Set(ctx, id, token, 10*time.Minute).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetUserToken(id string) (string, error) {
	rdb := database.ConnectRedis()
	ctx := context.Background()

	token, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}

func Verifyuser(id string) error {
	rdb := database.ConnectRedis()
	ctx := context.Background()
	conn, err := database.Connect()
	if err != nil {
		panic(err)
		return err
	}

	token, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return err
	}

	_, err2 := conn.Exec(context.Background(), "UPDATE users SET isverified = true WHERE email = $1", token)
	if err2 != nil {
		return err
	}

	_, err3 := rdb.Del(ctx, id).Result()
	if err3 != nil {
		return err
	}

	return nil
}

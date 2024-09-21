package database

import (
	"InventoryManagement/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, error) {

	conf := config.LoadConfig()

	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d",
			conf.Mongo.DbUser, conf.Mongo.DbPass, conf.Mongo.Host, conf.Mongo.Port),
	)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
		return nil, err
	}

	return client, nil
}

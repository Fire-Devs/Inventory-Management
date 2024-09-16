package database

import (
	"InventoryManagement/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {

	conf := config.LoadConfig()

	conn, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			conf.Database.DbUser,
			conf.Database.DbPass,
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.DbName,
		),
	)

	if err != nil {
		panic(err)
		return nil, err
	}

	return conn, nil

}

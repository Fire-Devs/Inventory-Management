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
			conf.Postgres.DbUser,
			conf.Postgres.DbPass,
			conf.Postgres.Host,
			conf.Postgres.Port,
			conf.Postgres.DbName,
		),
	)

	if err != nil {
		panic(err)
		return nil, err
	}

	return conn, nil

}

package database

import (
	"InventoryManagement/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func ConnectRedis() *redis.Client {

	conf := config.LoadConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Redis.Host, strconv.FormatInt(conf.Redis.Port, 10)),
	})

	return rdb
}

package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Postgres struct {
	Host   string
	Port   int64
	DbUser string
	DbPass string
	DbName string
}

type Mongo struct {
	Host   string
	Port   int64
	DbUser string
	DbPass string
}

type Redis struct {
	Host string
	Port int64
}

type Jwt struct {
	Secret string
}

type Server struct {
	Port string
}

type Config struct {
	Postgres Postgres
	Mongo    Mongo
	Redis    Redis
	Jwt      Jwt
	Server   Server
}

func LoadConfig() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresPort, _ := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 64)
	mongoPort, _ := strconv.ParseInt(os.Getenv("MONGO_PORT"), 10, 64)
	redisPort, _ := strconv.ParseInt(os.Getenv("REDIS_PORT"), 10, 64)

	return Config{
		Postgres: Postgres{
			Host:   os.Getenv("POSTGRES_HOST"),
			Port:   postgresPort,
			DbUser: os.Getenv("POSTGRES_USER"),
			DbPass: os.Getenv("POSTGRES_PASS"),
			DbName: os.Getenv("POSTGRES_DBNAME"),
		},
		Mongo: Mongo{
			Host:   os.Getenv("MONGO_HOST"),
			Port:   mongoPort,
			DbUser: os.Getenv("MONGO_USER"),
			DbPass: os.Getenv("MONGO_PASS"),
		},
		Redis: Redis{
			Host: os.Getenv("REDIS_HOST"),
			Port: redisPort,
		},
		Jwt: Jwt{
			Secret: os.Getenv("JWT_SECRET"),
		},
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
		},
	}
}

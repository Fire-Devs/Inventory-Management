package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Postgres struct {
	Host   string `toml:"host"`
	Port   int64  `toml:"port"`
	DbUser string `toml:"user"`
	DbPass string `toml:"pass"`
	DbName string `toml:"dbname"`
}

type Mongo struct {
	Host   string `toml:"host"`
	Port   int64  `toml:"port"`
	DbUser string `toml:"user"`
	DbPass string `toml:"pass"`
}

type Redis struct {
	Host string `toml:"host"`
	Port int64  `toml:"port"`
}

type Jwt struct {
	Secret string `toml:"secret"`
}

type Server struct {
	Port string `toml:"port"`
}

type Config struct {
	Postgres Postgres `toml:"postgres"`
	Mongo    Mongo    `toml:"mongo"`
	Redis    Redis    `toml:"redis"`
	Jwt      Jwt      `toml:"jwt"`
	Server   Server   `toml:"server"`
}

func LoadConfig() Config {
	var conf Config
	_, err := toml.DecodeFile("config.toml", &conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

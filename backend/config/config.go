package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Database struct {
	Host   string `toml:"host"`
	Port   int64  `toml:"port"`
	DbUser string `toml:"user"`
	DbPass string `toml:"pass"`
	DbName string `toml:"dbname"`
}
type Jwt struct {
	Secret string `toml:"secret"`
}

type Server struct {
	Port string `toml:"port"`
}

type Config struct {
	Database Database `toml:"database"`
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

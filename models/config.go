package models

type Config struct {
	Listen   Listen
	Database Database
}

type Listen struct {
	Host string
	Port int
}

type Database struct {
	Type string
}

func NewConfig() Config {
	return Config{
		Listen: Listen{
			Host: "0.0.0.0",
			Port: 8080,
		},
		Database: Database{
			Type: "sqlite",
		},
	}
}

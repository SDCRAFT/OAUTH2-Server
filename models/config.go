package models

type Config struct {
	Listen   Listen
	Database Database
	Captcha  Captcha
}

type Listen struct {
	Host string
	Port int
}

type Database struct {
	Account     Account
	Database    string
	Type        string
	TablePrefix string
	Host        string
	Port        int
	Paramters   map[string]string
}

type Captcha struct {
	Bulitin bool
}

type Account struct {
	Username string
	Password string
}

func NewConfig() Config {
	return Config{
		Listen: Listen{
			Host: "127.0.0.1",
			Port: 8080,
		},
		Database: Database{
			Account: Account{
				Username: "root",
				Password: "pa55w0rd",
			},
			Database:    "data.db",
			Type:        "sqlite",
			TablePrefix: "o2_",
			Host:        "127.0.0.1",
			Port:        3306,
			Paramters: map[string]string{
				"charset":   "utf8mb4",
				"parseTime": "true",
				"loc":       "UTC",
			},
		},
		Captcha: Captcha{
			Bulitin: true,
		},
	}
}

package database

import (
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	Models "sdcraft.fun/oauth2/models"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Failed to connect database: %v", err)
	}
	DB = db
	DB.AutoMigrate(&Models.User{})
}

package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	Models "sdcraft.fun/oauth2/models"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database, caused by: %v", err)
	}
	DB = db
	DB.AutoMigrate(&Models.User{})
}

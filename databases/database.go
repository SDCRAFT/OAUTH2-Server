package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func init() {
	_, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

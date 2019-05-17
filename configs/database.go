package configs

import (
	"github.com/shellbear/go-graphql-example/models"

	"github.com/jinzhu/gorm"
)

// DB is the database instance
var DB *gorm.DB

// InitDatabase init and return the database instance
func InitDatabase() {
	var err error
	DB, err = gorm.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Todo{})
}

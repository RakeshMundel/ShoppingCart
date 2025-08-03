package database

import (
	"fmt"
	"log"
	"shoppingcart/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("shopping.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

		err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Auto migration failed:", err)
	}

	fmt.Println(" Connected to database & migrated")
}

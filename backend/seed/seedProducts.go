package seed

import (
	"encoding/json"
	"fmt"
	"os"
	"shoppingcart/database"
	"shoppingcart/models"
)

func RunSeeder() {
	database.ConnectDatabase()

	file, err := os.ReadFile("products_seed.json")
	if err != nil {
		panic("Failed to read JSON file: " + err.Error())
	}

	var products []models.Product
	err = json.Unmarshal(file, &products)
	if err != nil {
		panic("Failed to unmarshal JSON: " + err.Error())
	}

	for _, product := range products {
		result := database.DB.Create(&product)
		if result.Error != nil {
			fmt.Printf("Failed to insert %s: %v\n", product.Name, result.Error)
		} else {
			fmt.Printf("✅ Inserted: %s\n", product.Name)
		}
	}
}


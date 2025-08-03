package controllers

import (
	"net/http"
	"shoppingcart/database"
	"shoppingcart/models"
	"fmt" // <--- Add this import statement

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProductsByCategory fetches products by a given category
func GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	var products []models.Product

	fmt.Println("Attempting to fetch products for category:", category) // <-- This line will now work

	result := database.DB.Where("category = ?", category).Find(&products)

	if result.Error != nil {
		fmt.Println("Error fetching products:", result.Error) // <-- This line will now work
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products for this category"})
		return
	}

	fmt.Println("Found", len(products), "products for category:", category) // <-- This line will now work
	c.JSON(http.StatusOK, products)
}

// You might also have GetAllProducts here, make sure it's present if needed
func GetAllProducts(c *gin.Context) {
	var products []models.Product

	result := database.DB.Find(&products)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
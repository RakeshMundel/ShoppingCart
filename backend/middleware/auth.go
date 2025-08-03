package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shoppingcart/database"
	"shoppingcart/models"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		var user models.User
		if err := database.DB.Where("token = ?", token).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		
		c.Set("user", user)
		c.Next()
	}
}

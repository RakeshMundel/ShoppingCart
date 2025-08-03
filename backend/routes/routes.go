package routes

import (
	"github.com/gin-gonic/gin"
	"shoppingcart/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)
	r.POST("/products", controllers.CreateProduct)
	//r.GET("/products", controllers.GetAllProducts)
    r.GET("/products/category/:category", controllers.GetProductsByCategory)
}


package routes

import (
    "github.com/gin-gonic/gin"
    "shoppingcart/controllers"
    "shoppingcart/middlewares" // ✅ Correct alias
)

func SetupRoutes(r *gin.Engine) {
    // Public Routes
    r.POST("/users", controllers.RegisterUser)
    r.POST("/users/login", controllers.LoginUser)
    r.GET("/users", controllers.ListUsers)

    r.POST("/items", controllers.CreateItem)
    r.GET("/items", controllers.ListItems)

    // Protected Routes
    authorized := r.Group("/", middlewares.AuthMiddleware())
    {
        authorized.POST("/carts", controllers.)
        authorized.GET("/carts", controllers.GetCart)

        authorized.POST("/orders", controllers.PlaceOrder)
        authorized.GET("/orders", controllers.GetOrders)
    }
}
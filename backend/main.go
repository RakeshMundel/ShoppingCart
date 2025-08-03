package main

import (
	"github.com/gin-gonic/gin"
	"shoppingcart/routes"
	"shoppingcart/database"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	routes.SetupRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
})


	r.Run(":4000")
}

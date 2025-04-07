package main

import (
	database "github.com/AVtheking/ticketo/db"
	routes "github.com/AVtheking/ticketo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	apiRoutes := router.Group("/api/v1")
	{
		routes.TheaterRoutes(db, apiRoutes)
	}

	router.Run(":8080")

}

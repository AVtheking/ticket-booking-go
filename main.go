package main

import (
	database "github.com/AVtheking/ticketo/db"
	"github.com/AVtheking/ticketo/routers"
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
		routers.CinemaRouter(apiRoutes)
	}

	router.Run(":8080")

}

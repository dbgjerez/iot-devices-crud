package main

import (
	"iot-devices-crud/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		health := new(controllers.HealthController)

		v1.GET("/health", health.Default)
		v1.GET("/device/:id", controllers.FindDevice)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	router.Run(":8080")
}

package main

import (
	"github.com/dbgjerez/iot-devices-crud/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		health := new(controllers.HealthController)
		deviceController := new(controllers.DeviceController)

		v1.GET("/health", health.Default)
		v1.POST("/device", deviceController.CreateDevice)
		v1.GET("/device/:id", deviceController.FindDevice)
		v1.GET("/device", deviceController.FindDevices)
		v1.DELETE("/device/:id", deviceController.DeleteDevice)
		v1.PUT("/device/:id", deviceController.UpdateDevice)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	router.Run(":8080")
}

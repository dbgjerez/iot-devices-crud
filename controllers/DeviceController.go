package controllers

import (
	"iot-devices-crud/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceController struct{}

func (d *DeviceController) FindDevice(c *gin.Context) {
	var device domain.Device

	c.JSON(http.StatusOK, gin.H{"data": device})
}

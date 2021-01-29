package controllers

import (
	"iot-devices-crud/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceController struct {
}

var repository = new(domain.DeviceRepository)

func (d *DeviceController) FindDevice(c *gin.Context) {
	device := repository.FindById(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": device})
}

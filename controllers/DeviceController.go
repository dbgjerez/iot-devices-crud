package controllers

import (
	"iot-devices-crud/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller constants
const (
	ParamPageName = "page" // pagination
	ParamSizeName = "size" // size string
	ParamDataName = "data" // data name
	ParamIDName   = "id"   // find by id

	PageDefaultValue = "1"  // pagination
	SizeDefaultValue = "10" // size string
	BaseDefault      = 10   // base to parse string to int
	BitSizeDefault   = 64   // default size to int when parse
)

type DeviceController struct {
}

var repository = new(domain.DeviceRepository)

func (d *DeviceController) CreateDevice(c *gin.Context) {
	var data domain.Device
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}
	device := repository.CreateDevice(data)
	c.JSON(http.StatusCreated, gin.H{ParamDataName: device})
}

func (d *DeviceController) FindDevice(c *gin.Context) {
	device := repository.FindById(c.Param(ParamIDName))
	c.JSON(http.StatusOK, gin.H{ParamDataName: device})
}

func (d *DeviceController) FindDevices(c *gin.Context) {
	page, err := strconv.ParseInt(c.DefaultQuery(ParamPageName, PageDefaultValue), BaseDefault, BitSizeDefault)
	size, err := strconv.ParseInt(c.DefaultQuery(ParamSizeName, SizeDefaultValue), BaseDefault, BitSizeDefault)
	if err != nil {
		panic(err)
	}
	devices := repository.FindAll(page, size)
	c.JSON(http.StatusOK, gin.H{ParamDataName: devices, ParamPageName: page, ParamSizeName: size})
}

func (d *DeviceController) DeleteDevice(c *gin.Context) {
	repository.DeleteDevice(c.Param(ParamIDName))
	c.JSON(http.StatusOK, gin.H{})
}

func (d *DeviceController) UpdateDevice(c *gin.Context) {
	var data domain.Device
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}
	repository.UpdateDevice(c.Param(ParamIDName), data)
	c.JSON(http.StatusOK, gin.H{})
}

package controllers

import (
	"iot-devices-crud/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller constants
const (
	ParamPageName = "page"     // pagination
	ParamSizeName = "size"     // size string
	ParamDataName = "data"     // data name
	ParamIDName   = "id"       // find by id
	ParamErrorMsg = "errorMsg" // error

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
		c.JSON(406, gin.H{ParamErrorMsg: "Provide relevant fields"})
		c.Abort()
		return
	}
	device, err := repository.CreateDevice(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ParamErrorMsg: err})
	} else {
		c.JSON(http.StatusCreated, gin.H{ParamDataName: device})
	}
}

func (d *DeviceController) FindDevice(c *gin.Context) {
	device, err := repository.FindById(c.Param(ParamIDName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ParamErrorMsg: err})
	} else if device == nil {
		c.JSON(http.StatusNotFound, gin.H{ParamDataName: device})
	} else {
		c.JSON(http.StatusOK, gin.H{ParamDataName: device})
	}
}

func (d *DeviceController) FindDevices(c *gin.Context) {
	page, pageErr := strconv.ParseInt(c.DefaultQuery(ParamPageName, PageDefaultValue), BaseDefault, BitSizeDefault)
	size, sizeErr := strconv.ParseInt(c.DefaultQuery(ParamSizeName, SizeDefaultValue), BaseDefault, BitSizeDefault)
	if pageErr != nil || sizeErr != nil {
		if pageErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{ParamErrorMsg: pageErr})
			c.Abort()
		}
		if sizeErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{ParamErrorMsg: sizeErr})
			c.Abort()
		}
	} else {
		devices, err := repository.FindAll(page, size)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ParamErrorMsg: err})
		} else {
			c.JSON(http.StatusOK, gin.H{ParamDataName: devices, ParamPageName: page, ParamSizeName: size})
		}
	}
}

func (d *DeviceController) DeleteDevice(c *gin.Context) {
	err := repository.DeleteDevice(c.Param(ParamIDName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ParamErrorMsg: err})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

func (d *DeviceController) UpdateDevice(c *gin.Context) {
	var data domain.Device
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{ParamErrorMsg: "Provide relevant fields"})
		c.Abort()
		return
	}
	repository.UpdateDevice(c.Param(ParamIDName), data)
	c.JSON(http.StatusOK, gin.H{})
}

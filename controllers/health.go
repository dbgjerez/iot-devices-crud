package controllers

import "github.com/gin-gonic/gin"

type HealthController struct{}

func (h *HealthController) Default(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}

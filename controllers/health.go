package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	status := "Health ok"
	c.JSON(http.StatusOK, gin.H{"data": status})
}

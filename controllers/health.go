package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	status := "social network service is up and running!"
	c.JSON(http.StatusOK, gin.H{"data": status})
}

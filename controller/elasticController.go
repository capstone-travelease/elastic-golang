package controller

import (
	"net/http"

	"github.com/billzayy/elastic-golang/services"
	"github.com/gin-gonic/gin"
)

func AllLocationController(c *gin.Context) {
	result, err := services.AllLocation()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
	}
	c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": result})
}

func SearchElasticController(c *gin.Context) {
	inputName := c.Query("name")

	result, err := services.SearchDoc(inputName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
	}
	c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": result})
}

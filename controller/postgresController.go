package controller

import (
	"net/http"

	"github.com/billzayy/elastic-golang/repositories"
	"github.com/gin-gonic/gin"
)

func PostgresControllerConnect(c *gin.Context) {
	result, err := repositories.ConnectDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
	}
	c.JSON(http.StatusAccepted, gin.H{"status": "Successful", "data": result})
}

package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var s Service

	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := CreateService(s); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create service"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "service created"})
}

func List(c *gin.Context) {
	data, err := GetServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch services"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}

	found, err := DeleteService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete service"})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "service deleted"})
}
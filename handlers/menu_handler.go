package handlers

import (
	"bookstore/config"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	var menu []models.MenuItem
	query := config.DB

	if restID := c.Query("restaurant_id"); restID != "" {
		query = query.Where("restaurant_id = ?", restID)
	}

	query.Find(&menu)
	c.JSON(http.StatusOK, menu)
}

func CreateMenuItem(c *gin.Context) {
	var newItem models.MenuItem
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&newItem)
	c.JSON(http.StatusCreated, newItem)
}

func DeleteMenuItem(c *gin.Context) {
	var item models.MenuItem
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu item not found"})
		return
	}
	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Menu item deleted"})
}

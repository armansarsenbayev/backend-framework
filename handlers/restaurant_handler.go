package handlers

import (
	"bookstore/config"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurants(c *gin.Context) {
	var restaurants []models.Restaurant
	config.DB.Find(&restaurants)
	c.JSON(http.StatusOK, restaurants)
}

func GetRestaurantByID(c *gin.Context) {
	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}
	c.JSON(http.StatusOK, restaurant)
}

func CreateRestaurant(c *gin.Context) {
	var newRestaurant models.Restaurant
	if err := c.ShouldBindJSON(&newRestaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&newRestaurant)
	c.JSON(http.StatusCreated, newRestaurant)
}

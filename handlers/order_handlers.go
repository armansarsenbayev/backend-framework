package handlers

import (
	"bookstore/config"
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var newOrder models.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder.Status = "pending"
	config.DB.Create(&newOrder)
	c.JSON(http.StatusCreated, newOrder)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	config.DB.Limit(limit).Offset(offset).Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {
	var order models.Order
	if err := config.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrderStatus(c *gin.Context) {
	var order models.Order
	if err := config.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var updateData struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&order).Update("status", updateData.Status)
	c.JSON(http.StatusOK, order)
}

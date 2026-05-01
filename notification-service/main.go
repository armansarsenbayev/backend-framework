package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type NotificationRequest struct {
	OrderID uint   `json:"order_id"`
	Status  string `json:"status"`
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		log.Printf("Incoming %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	r.POST("/notify", func(c *gin.Context) {
		var req NotificationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		fmt.Printf("[Notification Service] Received notification for Order #%d with status: %s\n", req.OrderID, req.Status)

		c.JSON(200, gin.H{
			"message": "Notification received successfully",
			"success": true,
		})
	})

	fmt.Println("NotificationService running on :8081")
	r.Run(":8081")
}

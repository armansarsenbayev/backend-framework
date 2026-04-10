package main

import (
	"bookstore/config"
	"bookstore/handlers"
	"bookstore/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Restaurant{}, &models.MenuItem{}, &models.Order{}, &models.User{})
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"} // Стандартный порт Vite React
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	router.GET("/restaurants", handlers.GetRestaurants)
	router.GET("/restaurants/:id", handlers.GetRestaurantByID)

	router.POST("/restaurants", handlers.AuthMiddleware(), handlers.CreateRestaurant)
	router.GET("/menu", handlers.GetMenu)
	router.POST("/menu", handlers.AuthMiddleware(), handlers.CreateMenuItem)
	router.DELETE("/menu/:id", handlers.DeleteMenuItem)

	router.POST("/orders", handlers.AuthMiddleware(), handlers.CreateOrder)
	router.GET("/orders", handlers.GetOrders)
	router.GET("/orders/:id", handlers.GetOrderByID)
	router.PUT("/orders/:id/status", handlers.UpdateOrderStatus)

	router.Run(":8080")
}

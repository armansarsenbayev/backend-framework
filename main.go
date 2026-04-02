package main

import (
	"bookstore/config"
	"bookstore/handlers"
	"bookstore/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Book{}, &models.Author{}, &models.Category{})

	router := gin.Default()


	router.GET("/authors", handlers.GetAuthors)
	router.GET("/authors/:id", handlers.GetAuthorByID)
	router.POST("/authors", handlers.CreateAuthor)

	router.GET("/categories", handlers.GetCategories)
	router.POST("/categories", handlers.CreateCategory)

	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBookByID)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	router.Run(":8080")
}
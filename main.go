package main

import (
    "bookstore/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/authors", handlers.GetAuthors)
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
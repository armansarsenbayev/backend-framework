package handlers

import (
	"bookstore/config"
	"bookstore/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	config.DB.Find(&authors) 
	c.JSON(http.StatusOK, authors)
}

func GetAuthorByID(c *gin.Context) {
	var author models.Author
	if err := config.DB.First(&author, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

func CreateAuthor(c *gin.Context) {
	var newAuthor models.Author
	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&newAuthor) 
	c.JSON(http.StatusCreated, newAuthor)
}
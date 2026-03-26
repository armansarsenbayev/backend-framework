package handlers

import (
    "bookstore/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

var authors = []models.Author{}
var nextAuthorID = 1

func GetAuthors(c *gin.Context) {
    c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
    var newAuthor models.Author
    
    if err := c.ShouldBindJSON(&newAuthor); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newAuthor.ID = nextAuthorID
    nextAuthorID++
    authors = append(authors, newAuthor)

    c.JSON(http.StatusCreated, newAuthor)
}
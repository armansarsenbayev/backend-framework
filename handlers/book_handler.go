package handlers

import (
    "bookstore/models"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

var books = []models.Book{}
var nextBookID = 1

func GetBooks(c *gin.Context) {
    filteredBooks := []models.Book{}

    categoryParam := c.Query("category_id")
    if categoryParam != "" {
        catID, _ := strconv.Atoi(categoryParam)
        for _, b := range books {
            if b.CategoryID == catID {
                filteredBooks = append(filteredBooks, b)
            }
        }
    } else {
        filteredBooks = books
    }

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    start := (page - 1) * limit
    end := start + limit

    if start >= len(filteredBooks) {
        c.JSON(http.StatusOK, []models.Book{})
        return
    }
    if end > len(filteredBooks) {
        end = len(filteredBooks)
    }

    c.JSON(http.StatusOK, filteredBooks[start:end])
}

func CreateBook(c *gin.Context) {
    var newBook models.Book
    if err := c.ShouldBindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newBook.ID = nextBookID
    nextBookID++
    books = append(books, newBook)

    c.JSON(http.StatusCreated, newBook)
}

func GetBookByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    for _, b := range books {
        if b.ID == id {
            c.JSON(http.StatusOK, b)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func UpdateBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    var updatedData models.Book
    if err := c.ShouldBindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, b := range books {
        if b.ID == id {
            updatedData.ID = id
            books[i] = updatedData
            c.JSON(http.StatusOK, updatedData)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func DeleteBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    for i, b := range books {
        if b.ID == id {
            books = append(books[:i], books[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
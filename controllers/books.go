package controllers

import (
	"net/http"

	"github.com/felipefbs/bookShelf/models"
	"github.com/gin-gonic/gin"
)

// CreateBookInput struct
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookInput struct
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GetAllBooks function returns all books from the database
func GetAllBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook function adds a new book to the database
func CreateBook(c *gin.Context) {
	var input CreateBookInput

	//Validation of the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GetBook function returns a single book from the database based on it ID
func GetBook(c *gin.Context) {
	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook changes information of a existing book
func UpdateBook(c *gin.Context) {
	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

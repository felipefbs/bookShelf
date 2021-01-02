package controllers

import (
	"net/http"

	"github.com/felipefbs/bookShelf/models"
	"github.com/gin-gonic/gin"
)

// GetAllBooks function
func GetAllBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

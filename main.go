package main

import (
	"github.com/felipefbs/bookShelf/controllers"
	"github.com/felipefbs/bookShelf/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", controllers.GetAllBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.GetBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run()
}

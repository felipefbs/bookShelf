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

	router.Run()
}

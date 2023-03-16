package main

import (
	"challenges-three/controllers"
	"challenges-three/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.GetConnection()
	defer db.Close()
	route := gin.Default()

	route.POST("/books", controllers.AddBook(db))
	route.GET("/books", controllers.GetAllBook(db))
	route.GET("/books/:bookID", controllers.GetBook(db))
	route.PUT("/books/:bookID", controllers.UpdateBook(db))
	route.DELETE("/books/:bookID", controllers.DeleteBook(db))

	route.Run(":8080")
}

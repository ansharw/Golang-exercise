package routers

import (
	"challenges-two/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine  {
	router := gin.Default()

	router.POST("/books", controllers.AddBook)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.GET("/books/:bookID", controllers.GetBookById)
	router.GET("/books", controllers.GetAllBook)
	router.DELETE("books/:bookID", controllers.DeleteBook)
	return router
}
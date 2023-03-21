package routers

import (
	"showcase/controllers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Route(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	route.POST("/books", controllers.AddBook(db))
	route.GET("/books", controllers.GetAllBook(db))
	route.GET("/books/:bookID", controllers.GetBook(db))
	route.PUT("/books/:bookID", controllers.UpdateBook(db))
	route.DELETE("/books/:bookID", controllers.DeleteBook(db))

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return route
}

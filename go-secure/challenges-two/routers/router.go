package routers

import (
	"challenges-two/controllers"
	"challenges-two/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorizations(), controllers.UpdateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorizations(), controllers.GetProduct)
		productRouter.GET("/", middlewares.ProductAuthorizations(), controllers.GetAllProduct)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorizations(), controllers.DeleteProduct)
	}

	return r
}

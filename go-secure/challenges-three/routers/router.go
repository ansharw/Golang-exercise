package routers

import (
	"challenges-three/controllers"
	"challenges-three/database"
	"challenges-three/middlewares"
	"challenges-three/repository"
	"challenges-three/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	db := database.GetConnection()
	validate := validator.New()

	repoProduct := repository.NewProductRepository()

	serviceProduct := services.NewProductService(db, repoProduct, *validate)

	handlerProduct := controllers.NewProductHandler(serviceProduct)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		// productRouter.POST("/", controllers.CreateProduct)
		// productRouter.PUT("/:productId", middlewares.ProductAuthorizations(), controllers.UpdateProduct)
		// productRouter.GET("/:productId", middlewares.ProductAuthorizations(), controllers.GetProduct)
		productRouter.GET("/", middlewares.ProductAuthorizations(), handlerProduct.GetAllProducts)
		// productRouter.DELETE("/:productId", middlewares.ProductAuthorizations(), controllers.DeleteProduct)
	}

	return r
}

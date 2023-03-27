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
	repoUser := repository.NewUserRepository()

	serviceProduct := services.NewProductService(db, repoProduct, *validate)
	serviceUser := services.NewUserService(db, repoUser, *validate)

	handlerProduct := controllers.NewProductHandler(serviceProduct)
	handlerUser := controllers.NewUserHandler(serviceUser)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", handlerUser.Register)
		userRouter.POST("/login", handlerUser.Login)
	}
	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", handlerProduct.CreateProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorizations(), handlerProduct.UpdateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorizations(), handlerProduct.GetProduct)
		productRouter.GET("/", middlewares.ProductAuthorizations(), handlerProduct.GetAllProducts)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorizations(), handlerProduct.DeleteProduct)
	}

	return r
}

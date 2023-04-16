package routers

import (
	"MyGram/controllers"
	"MyGram/database"
	"MyGram/middlewares"
	"MyGram/model"
	"MyGram/repository"
	"MyGram/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	_ "MyGram/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title          	Swagger MyGram API
// @version        	1.0
// @description    	This is a sample server MyGram server.
// @contact.name   	ansharw
// @host      		localhost:8080
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description description: Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345".
// @externalDocs.description  OpenAPI
func StartApp() *gin.Engine {
	r := gin.Default()
	db := database.GetConnection()
	validate := validator.New()

	repoSocialMedia := repository.NewSocialMediaRepository()
	repoPhoto := repository.NewPhotoRepository()
	repoComment := repository.NewCommentRepository()
	repoUser := repository.NewUserRepository()

	serviceSocialMedia := services.NewSocialMediaService(db, repoSocialMedia)
	servicePhoto := services.NewPhotoService(db, repoPhoto)
	serviceComment := services.NewCommentService(db, repoComment)
	serviceUser := services.NewUserService(db, repoUser)

	handlerSocialMedia := controllers.NewSocialMediaHandler(serviceSocialMedia, *validate)
	handlerPhoto := controllers.NewPhotoHandler(servicePhoto, *validate)
	handlerComment := controllers.NewCommentHandler(serviceComment, *validate)
	handlerUser := controllers.NewUserHandler(serviceUser, *validate)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", handlerUser.Register)
		userRouter.POST("/login", handlerUser.Login)
	}

	socialMediaRouter := r.Group("/socialmedia")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", handlerSocialMedia.CreateSocialMedia)
		socialMediaRouter.GET("/", handlerSocialMedia.GetAllSocialMedia)
		socialMediaRouter.GET("/:socialMediaId", middlewares.SocialMediaAuthorizations(&model.SocialMedia{}), handlerSocialMedia.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorizations(&model.SocialMedia{}), handlerSocialMedia.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorizations(&model.SocialMedia{}), handlerSocialMedia.DeleteSocialMedia)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", handlerPhoto.CreatePhoto)
		photoRouter.GET("/", handlerPhoto.GetAllPhoto)
		photoRouter.GET("/:photoId", middlewares.PhotoAuthorizations(&model.Photo{}), handlerPhoto.GetPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorizations(&model.Photo{}), handlerPhoto.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorizations(&model.Photo{}), handlerPhoto.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", handlerComment.CreateComment)
		commentRouter.GET("/", handlerComment.GetAllComment)
		commentRouter.GET("/:commentId", middlewares.CommentAuthorizations(&model.Comments{}), handlerComment.GetComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorizations(&model.Comments{}), handlerComment.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorizations(&model.Comments{}), handlerComment.DeleteComment)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

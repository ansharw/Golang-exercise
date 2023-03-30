package routers

import (
	"MyGram/controllers"
	"MyGram/database"
	"MyGram/middlewares"
	"MyGram/repository"
	"MyGram/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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
		socialMediaRouter.GET("/", middlewares.SocialMediaAuthorizations(), handlerSocialMedia.GetAllSocialMedia)
		socialMediaRouter.GET("/:socialMediaId", middlewares.SocialMediaAuthorizations(), handlerSocialMedia.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorizations(), handlerSocialMedia.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorizations(), handlerSocialMedia.DeleteSocialMedia)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", handlerPhoto.CreatePhoto)
		photoRouter.GET("/", middlewares.PhotoAuthorizations(), handlerPhoto.GetAllPhoto)
		photoRouter.GET("/:photoId", middlewares.PhotoAuthorizations(), handlerPhoto.GetPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorizations(), handlerPhoto.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorizations(), handlerPhoto.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", handlerComment.CreateComment)
		commentRouter.GET("/", middlewares.CommentAuthorizations(), handlerComment.GetAllComment)
		commentRouter.GET("/:commentId", middlewares.CommentAuthorizations(), handlerComment.GetComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorizations(), handlerComment.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorizations(), handlerComment.DeleteComment)
	}

	return r
}

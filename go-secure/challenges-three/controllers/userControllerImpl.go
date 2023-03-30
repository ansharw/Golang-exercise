package controllers

import (
	"challenges-three/helpers"
	"challenges-three/models"
	"challenges-three/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

var (
	appJson = "application/json"
)

func (handler *userHandler) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := models.User{}

	if contentType == appJson {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	if token, err := handler.userService.Login(c, User); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	} else {
		// token := helpers.GenerateToken(User.ID, User.Email)

		c.JSON(http.StatusCreated, gin.H{
			"token": token,
		})
	}
}

func (handler *userHandler) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := models.User{}

	if contentType == appJson {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	if user, err := handler.userService.Register(c, User); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Email Already Exists/Password Invalid",
			"message": "The email address you entered already exists",
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id":        user.ID,
			"email":     user.Email,
			"full_name": user.Fullname,
		})
	}
}

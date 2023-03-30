package controllers

import (
	"MyGram/helpers"
	"MyGram/model"
	"MyGram/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	appJson = "application/json"
)

type userHandler struct {
	userService services.UserService
	validate    *validator.Validate
}

func NewUserHandler(userService services.UserService, validator_ validator.Validate) *userHandler {
	return &userHandler{
		userService: userService,
		validate:    &validator_,
	}
}

// Login
// User can login account
func (handler *userHandler) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := model.RequestUserLogin{}

	// binding email, password
	if contentType == appJson {
		if err := c.ShouldBindJSON(&User); err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				var errMsg string
				for _, e := range errors {
					switch e.Field() {
					case "Email":
						errMsg = "Invalid email."
					case "Password":
						errMsg = "Invalid password."
					}
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request json",
					"message": errMsg,
				})
				return
			}
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				var errMsg string
				for _, e := range errors {
					switch e.Field() {
					case "Email":
						errMsg = "Invalid email."
					case "Password":
						errMsg = "Invalid password."
					}
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request json",
					"message": errMsg,
				})
				return
			}
		}
	}

	if token, err := handler.userService.Login(c, User); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"token": token,
		})
	}
}

// Register user
// User can register account
func (handler *userHandler) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := model.User{}

	// bind username, email, password, age
	if contentType == appJson {
		if err := c.ShouldBindJSON(&User); err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				var errMsg string
				for _, e := range errors {
					switch e.Field() {
					case "Username":
						errMsg = "Invalid username."
					case "Password":
						errMsg = "Invalid password."
					case "Email":
						errMsg = "Invalid email"
					case "Age":
						errMsg = "Invalid age."
					}
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request json",
					"message": errMsg,
				})
				return
			}
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				var errMsg string
				for _, e := range errors {
					switch e.Field() {
					case "Username":
						errMsg = "Invalid username."
					case "Password":
						errMsg = "Invalid password."
					case "Email":
						errMsg = "Invalid email"
					case "Age":
						errMsg = "Invalid age."
					}
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request form",
					"message": errMsg,
				})
				return
			}
		}
	}

	if user, err := handler.userService.Register(c, User); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Email Already Exists/Password Invalid",
			"message": "The email address you entered already exists/password invalid",
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"username": user.Username,
		})
	}
}

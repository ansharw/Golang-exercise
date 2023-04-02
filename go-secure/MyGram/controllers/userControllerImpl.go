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
// Login godoc
// @Summary Login user
// @Description Login user
// @Tags User
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Param requestLogin body model.RequestUserLogin true "login user"
// @Success 201 {object} model.ResponseErrorGeneral
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 401 {object} model.ResponseErrorGeneral
// @Router /users/login [post]
func (handler *userHandler) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := model.RequestUserLogin{}

	// binding email, password
	var err error
	if contentType == appJson {
		err = c.ShouldBindJSON(&User)
	} else {
		err = c.ShouldBind(&User)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(User); err != nil {
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
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
				Status:  "Bad Request json",
				Message: errMsg,
			})
			return
		}
	}

	// if contentType == appJson {
	// 	if err := c.ShouldBindJSON(&User); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Email":
	// 					errMsg = "Invalid email."
	// 				case "Password":
	// 					errMsg = "Invalid password."
	// 				}
	// 			}
	// 			c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 				Status:  "Bad Request json",
	// 				Message: errMsg,
	// 			})
	// 			return
	// 		}
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&User); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Email":
	// 					errMsg = "Invalid email."
	// 				case "Password":
	// 					errMsg = "Invalid password."
	// 				}
	// 			}
	// 			c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 				Status:  "Bad Request form",
	// 				Message: errMsg,
	// 			})
	// 			return
	// 		}
	// 	}
	// }

	if token, err := handler.userService.Login(c, User); err != nil {
		c.JSON(http.StatusUnauthorized, model.ResponseErrorGeneral{
			Status:  "Unauthorized",
			Message: "invalid email/password",
		})
		return
	} else {
		c.JSON(http.StatusCreated, model.ResponseToken{
			Token: token,
		})
	}
}

// Register user
// User can register account
// Register godoc
// @Summary Register user
// @Description Register user
// @Tags User
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Param requestRegister body model.RequestUserRegister true "Register user"
// @Success 201 {object} model.ResponseErrorGeneral
// @Failure 400 {object} model.ResponseErrorGeneral
// @Router /users/register [post]
func (handler *userHandler) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := model.RequestUserRegister{}

	// bind username, email, password, age
	var err error
	if contentType == appJson {
		err = c.ShouldBindJSON(&User)
	} else {
		err = c.ShouldBind(&User)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(User); err != nil {
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
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
				Status:  "Bad Request json/form",
				Message: errMsg,
			})
			return
		}
	}

	// if contentType == appJson {
	// 	if err := c.ShouldBindJSON(&User); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Username":
	// 					errMsg = "Invalid username."
	// 				case "Password":
	// 					errMsg = "Invalid password."
	// 				case "Email":
	// 					errMsg = "Invalid email"
	// 				case "Age":
	// 					errMsg = "Invalid age."
	// 				}
	// 			}
	// 			c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 				Status:  "Bad Request json",
	// 				Message: errMsg,
	// 			})
	// 			return
	// 		}
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&User); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Username":
	// 					errMsg = "Invalid username."
	// 				case "Password":
	// 					errMsg = "Invalid password."
	// 				case "Email":
	// 					errMsg = "Invalid email"
	// 				case "Age":
	// 					errMsg = "Invalid age."
	// 				}
	// 			}
	// 			c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 				Status:  "Bad Request form",
	// 				Message: errMsg,
	// 			})
	// 			return
	// 		}
	// 	}
	// }

	if user, err := handler.userService.Register(c, User); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Email Already Exists/Password Invalid",
			Message: "The email address you entered already exists/password invalid",
		})
		return
	} else {
		c.JSON(http.StatusCreated, model.ResponseRegistered{
			Id:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		})
	}
}

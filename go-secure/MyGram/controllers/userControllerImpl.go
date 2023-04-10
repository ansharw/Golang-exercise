package controllers

import (
	"MyGram/helpers"
	"MyGram/model"
	"MyGram/services"
	"fmt"
	"net/http"
	"strings"

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

	// ----------------------- validation version 1
	// if err != nil {
	// 	if errors, ok := err.(validator.ValidationErrors); ok {
	// 		var errMsg string
	// 		for _, e := range errors {
	// 			switch {
	// 			case e.Field() == "Password":
	// 				if e.Tag() == "required" {
	// 					field, _ := reflect.TypeOf(User).FieldByName("Password")
	// 					errMsg = fmt.Sprintf("%s is required.", field.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "password cannot be empty"
	// 				}
	// 			case e.Field() == "Email":
	// 				if len(errors) == 2 {
	// 					field_email, _ := reflect.TypeOf(User).FieldByName("Email")
	// 					errMsg = fmt.Sprintf("%s is required and invalid email.", field_email.Tag.Get("json"))
	// 				} else if e.Tag() == "required" {
	// 					field_email, _ := reflect.TypeOf(User).FieldByName("Email")
	// 					errMsg = fmt.Sprintf("%s is required.", field_email.Tag.Get("json"))
	// 				} else if e.Tag() == "email" {
	// 					field_email, _ := reflect.TypeOf(User).FieldByName("Email")
	// 					errMsg = fmt.Sprintf("%s is invalid.", field_email.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "email cannot be empty"
	// 				}
	// 			default:
	// 				if e.Tag() == "required" {
	// 					field_email, _ := reflect.TypeOf(User).FieldByName("Email")
	// 					field_password, _ := reflect.TypeOf(User).FieldByName("Password")
	// 					errMsg = fmt.Sprintf("%s and %s is required.", field_email.Tag.Get("json"), field_password.Tag.Get("json"))
	// 				} else if e.Tag() == "required" || e.Tag() == "email" {
	// 					field_email, _ := reflect.TypeOf(User).FieldByName("Email")
	// 					field_password, _ := reflect.TypeOf(User).FieldByName("Password")
	// 					errMsg = fmt.Sprintf("%s is required and %s is invalid email.", field_password.Tag.Get("json"), field_email.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "email and password cannot be empty"
	// 				}
	// 			}
	// 		}
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 			Status:  "Bad Request json/form",
	// 			Message: errMsg,
	// 		})
	// 		return
	// 	} else {
	// 		// all error json / form
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 			Status:  "Bad Request json/form",
	// 			Message: err.Error(),
	// 		})
	// 		return
	// 	}
	// }

	// ----------------------- validation version 2
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// Convert validation errors to map of error messages
			errorsMap := make(map[string]string)
			for _, validationError := range validationErrors {
				// Use the field name as the error key
				field := validationError.Field()
				// validation error tag
				switch validationError.Tag() {
				case "email":
					switch field {
					case "Email":
						errorsMap[field] = fmt.Sprintf("%s is invalid", field)
					}
				default:
					errorsMap[field] = fmt.Sprintf("%s is required", field)
				}
			}
			// Join error messages into a single string
			var errorMessages []string
			for _, errorMessage := range errorsMap {
				errorMessages = append(errorMessages, errorMessage)
			}
			errorMessageString := strings.Join(errorMessages, ", ")

			// Return errors map as JSON response
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request json/form",
				"message": errorMessageString,
			})
			return
		}
		// Error json unmarshal / error binding json/form
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request json/form",
			"message": err.Error(),
		})
		return
	}

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
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// Convert validation errors to map of error messages
			errorsMap := make(map[string]string)
			for _, validationError := range validationErrors {
				// Use the field name as the error key
				field := validationError.Field()
				// validation error tag
				switch validationError.Tag() {
				case "unique":
					switch field {
					case "Username", "Email":
						errorsMap[field] = fmt.Sprintf("%s is must be unique", field)
					}
				case "min":
					switch field {
					case "Password":
						errorsMap[field] = fmt.Sprintf("%s must be at least 6 characters long", field)
					}
				case "gte":
					switch field {
					case "Age":
						errorsMap[field] = fmt.Sprintf("%s must be greater than or equal to 8", field)
					}
				case "email":
					switch field {
					case "Email":
						errorsMap[field] = fmt.Sprintf("%s is invalid", field)
					}
				default:
					errorsMap[field] = fmt.Sprintf("%s is required", field)
				}
			}
			// Join error messages into a single string
			var errorMessages []string
			for _, errorMessage := range errorsMap {
				errorMessages = append(errorMessages, errorMessage)
			}
			errorMessageString := strings.Join(errorMessages, ", ")

			// Return errors map as JSON response
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request json/form",
				"message": errorMessageString,
			})
			return
		}
		// Error json unmarshal / error binding json/form
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request json/form",
			"message": err.Error(),
		})
		return
	}

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

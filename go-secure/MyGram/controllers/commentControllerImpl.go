package controllers

import (
	"MyGram/helpers"
	"MyGram/model"
	"MyGram/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type commentHandler struct {
	commentService services.CommentService
	validate       *validator.Validate
}

func NewCommentHandler(commentService services.CommentService, validator validator.Validate) *commentHandler {
	return &commentHandler{
		commentService: commentService,
		validate:       &validator,
	}
}

// Get all comment by photo id
// all user can access to show all comment by photo id
// Get all comment by photo id godoc
// @Summary Get all comment by photo id user
// @Description Get all comment by photo id user
// @Tags Comment
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param requestGet body model.RequestGetComments true "Get All Comment By photo id"
// @Success 200 {array} model.Comments
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 401 {object} model.ResponseErrorGeneral
// @Failure 404 {object} model.ResponseErrorGeneral
// @Failure 500 {object} model.ResponseErrorGeneral
// @Router /comment [get]
func (handler *commentHandler) GetAllComment(c *gin.Context) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comments := model.RequestGetComments{}

	// bind photo_id
	var err error
	if contentType == appJson {
		err = c.ShouldBindJSON(&comments)
	} else {
		err = c.ShouldBind(&comments)
	}

	// ----------------------- validation version 1
	// if err != nil {
	// 	if errors, ok := err.(validator.ValidationErrors); ok {
	// 		var errMsg string
	// 		for _, e := range errors {
	// 			switch {
	// 			case e.Field() == "PhotoID":
	// 				if e.Tag() == "required" {
	// 					field, _ := reflect.TypeOf(comment).FieldByName("PhotoID")
	// 					errMsg = fmt.Sprintf("%s is required.", field.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "photo_id cannot be empty"
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
				case "required":
					switch field {
					case "PhotoID":
						errorsMap[field] = fmt.Sprintf("%s is required", field)
					}
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

	res, err := handler.commentService.FindAllByPhotoId(c, comments.PhotoID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseErrorGeneral{
			Status:  "Internal Server Error",
			Message: fmt.Sprintf("Error retrieving all comment user: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Get comment by photo id
// all user can access read comment by photo id
// Get comment by photo id godoc
// @Summary Get comment by photo id user
// @Description Get comment by photo id user
// @Tags Comment
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param commentId path int true "Comment ID"
// @Param requestGet body model.RequestGetComments true "Get Comment By photo id and comment id in path params"
// @Success 200 {object} model.Comments
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 401 {object} model.ResponseErrorGeneral
// @Failure 404 {object} model.ResponseErrorGeneral
// @Router /comment/{commentId} [get]
func (handler *commentHandler) GetComment(c *gin.Context) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestGetComments{}

	// Get param commentId
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil || uint(commentId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: "invalid comment ID",
		})
		return
	}

	// bind photo_id
	if contentType == appJson {
		err = c.ShouldBindJSON(&comment)
	} else {
		err = c.ShouldBind(&comment)
	}

	// ----------------------- validation version 1
	// if err != nil {
	// 	if errors, ok := err.(validator.ValidationErrors); ok {
	// 		var errMsg string
	// 		for _, e := range errors {
	// 			switch {
	// 			case e.Field() == "PhotoID":
	// 				if e.Tag() == "required" {
	// 					field, _ := reflect.TypeOf(comment).FieldByName("PhotoID")
	// 					errMsg = fmt.Sprintf("%s is required.", field.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "photo_id cannot be empty"
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
				case "required":
					switch field {
					case "PhotoID":
						errorsMap[field] = fmt.Sprintf("%s is required", field)
					}
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

	if res, err := handler.commentService.FindByPhotoId(c, comment.PhotoID, uint(commentId)); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.ResponseErrorGeneral{
			Status:  "Data Not Found",
			Message: fmt.Sprintf("Comment Data with id: %d not found\n", commentId),
		})
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Create comment by photo id and user id
// all user can create comment by photo id and user id
// Create comment photo godoc
// @Summary Create comment photo user
// @Description Create comment photo user
// @Tags Comment
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param requestCreate body model.RequestComments true "Create Comment photo"
// @Success 201 {object} model.Comments
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 401 {object} model.ResponseErrorGeneral
// @Router /comment [post]
func (handler *commentHandler) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestComments{}

	// bind message, photo_id
	var err error
	if contentType == appJson {
		err = c.ShouldBindJSON(&comment)
	} else {
		err = c.ShouldBind(&comment)
	}

	// ----------------------- validation version 1
	// if err != nil {
	// 	if errors, ok := err.(validator.ValidationErrors); ok {
	// 		var errMsg string
	// 		for _, e := range errors {
	// 			switch {
	// 			case e.Field() == "PhotoID":
	// 				if e.Tag() == "required" {
	// 					field, _ := reflect.TypeOf(comment).FieldByName("PhotoID")
	// 					errMsg = fmt.Sprintf("%s is required.", field.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "photo_id cannot be empty"
	// 				}
	// 			case e.Field() == "Message":
	// 				if e.Tag() == "required" {
	// 					field_message, _ := reflect.TypeOf(comment).FieldByName("Message")
	// 					errMsg = fmt.Sprintf("%s is required.", field_message.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "message cannot be empty"
	// 				}
	// 			default:
	// 				if e.Tag() == "required" {
	// 					field_message, _ := reflect.TypeOf(comment).FieldByName("Message")
	// 					field_photo_id, _ := reflect.TypeOf(comment).FieldByName("PhotoID")
	// 					errMsg = fmt.Sprintf("%s and %s is required.", field_message.Tag.Get("json"), field_photo_id.Tag.Get("json"))
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
				case "required":
					switch field {
					case "PhotoID", "Message":
						errorsMap[field] = fmt.Sprintf("%s is required", field)
					}
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

	if res, err := handler.commentService.Create(c, comment, userID, comment.PhotoID); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// Update comment by photo id and user id
// User can update comment by photo id and user id
// Update comment by photo id godoc
// @Summary Update comment by photo id user
// @Description Update comment by photo id user
// @Tags Comment
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param commentId path int true "Comment ID"
// @Param requestUpdate body model.RequestComments true "Update Comment by photo id user"
// @Success 200 {object} model.Comments
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 401 {object} model.ResponseErrorGeneral
// @Router /comment/{commentId} [put]
func (handler *commentHandler) UpdateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestComments{}

	// Get param commentId
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: "invalid comment ID",
		})
		return
	}

	// bind message, photo_id
	if contentType == appJson {
		err = c.ShouldBindJSON(&comment)
	} else {
		err = c.ShouldBind(&comment)
	}

	// ----------------------- validation version 1
	// if err != nil {
	// 	if errors, ok := err.(validator.ValidationErrors); ok {
	// 		var errMsg string
	// 		for _, e := range errors {
	// 			switch {
	// 			case e.Field() == "PhotoID":
	// 				if e.Tag() == "required" {
	// 					field, _ := reflect.TypeOf(comment).FieldByName("PhotoID")
	// 					errMsg = fmt.Sprintf("%s is required.", field.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "photo_id cannot be empty"
	// 				}
	// 			case e.Field() == "Message":
	// 				if e.Tag() == "required" {
	// 					field_message, _ := reflect.TypeOf(comment).FieldByName("Message")
	// 					errMsg = fmt.Sprintf("%s is required.", field_message.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "message cannot be empty"
	// 				}
	// 			default:
	// 				if e.Tag() == "required" {
	// 					field_message, _ := reflect.TypeOf(comment).FieldByName("Message")
	// 					field_photo_id, _ := reflect.TypeOf(comment).FieldByName("PhotoID")
	// 					errMsg = fmt.Sprintf("%s and %s is required.", field_message.Tag.Get("json"), field_photo_id.Tag.Get("json"))
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
				case "required":
					switch field {
					case "PhotoID", "Message":
						errorsMap[field] = fmt.Sprintf("%s is required", field)
					}
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

	if res, err := handler.commentService.Update(c, comment, uint(commentId), userID, comment.PhotoID); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Delete comment by photo id and user id
// User can delete comment by photo id and user id
// Delete comment by photo id godoc
// @Summary Delete comment by photo id user
// @Description Delete comment by photo id user
// @Tags Comment
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param commentId path int true "Comment ID"
// @Success 200 {object} model.ResponseDeleted
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 401 {object} model.ResponseErrorGeneral
// @Failure 500 {object} model.ResponseErrorGeneral
// @Router /comment/{commentId} [delete]
func (handler *commentHandler) DeleteComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestDeleteComments{}

	// Get param commentId
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil || uint(commentId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: "invalid comment ID",
		})
		return
	}

	// bind photo_id
	if contentType == appJson {
		err = c.ShouldBindJSON(&comment)
	} else {
		err = c.ShouldBind(&comment)
	}

	// ----------------------- validation version 1
	// if err != nil {
	// 	if errors, ok := err.(validator.ValidationErrors); ok {
	// 		var errMsg string
	// 		for _, e := range errors {
	// 			switch {
	// 			case e.Field() == "PhotoID":
	// 				if e.Tag() == "required" {
	// 					field, _ := reflect.TypeOf(comment).FieldByName("PhotoID")
	// 					errMsg = fmt.Sprintf("%s is required.", field.Tag.Get("json"))
	// 				} else {
	// 					errMsg = "photo_id cannot be empty"
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
				case "required":
					switch field {
					case "PhotoID":
						errorsMap[field] = fmt.Sprintf("%s is required", field)
					}
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

	if err := handler.commentService.Delete(c, uint(commentId), userID, comment.PhotoID); err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseErrorGeneral{
			Status:  "Internal Server Error",
			Message: "Failed to delete comment",
		})
		return
	} else {
		c.JSON(http.StatusOK, model.ResponseErrorGeneral{
			Message: "Data Comment deleted successfully",
		})
	}
}

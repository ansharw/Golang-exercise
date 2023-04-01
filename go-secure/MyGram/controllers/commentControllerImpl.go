package controllers

import (
	"MyGram/helpers"
	"MyGram/model"
	"MyGram/services"
	"fmt"
	"net/http"
	"strconv"

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
// @Produce json
// @Security token
// @securityDefinitions.apikey token
// @Success 200 {array} model.Comment
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 500 {object} model.ResponseErrorGeneral
// @Router /comment [get]
func (handler *commentHandler) GetAllComment(c *gin.Context) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.Comment{}

	// bind photo_id
	var err error
	if contentType == appJson {
		err = c.ShouldBindJSON(&comment)
	} else {
		err = c.ShouldBind(&comment)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(comment); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range errors {
				switch e.Field() {
				case "Message":
					errMsg = "Invalid title."
				case "Caption":
					errMsg = "Invalid caption."
				case "PhotoID":
					errMsg = "Invalid photo_id."
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
	// 	if err := c.ShouldBindJSON(&comment); err != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 			Status:  "Bad Request",
	// 			Message: err.Error(),
	// 		})
	// 		return
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&comment); err != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 			Status:  "Bad Request",
	// 			Message: err.Error(),
	// 		})
	// 		return
	// 	}
	// }

	res, err := handler.commentService.FindAllByPhotoId(c, comment.PhotoID)
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
// @Produce json
// @Security token
// @securityDefinitions.apikey token
// @Param commentId path int true "Comment ID"
// @Success 200 {object} model.Comment
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 404 {object} model.ResponseErrorGeneral
// @Router /comment/{photoId} [get]
func (handler *commentHandler) GetComment(c *gin.Context) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.Comment{}

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
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(comment); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range errors {
				switch e.Field() {
				case "Message":
					errMsg = "Invalid title."
				case "Caption":
					errMsg = "Invalid caption."
				case "PhotoID":
					errMsg = "Invalid photo_id."
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
	// 	if err := c.ShouldBindJSON(&comment); err != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 			Status:  "Bad Request",
	// 			Message: err.Error(),
	// 		})
	// 		return
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&comment); err != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 			Status:  "Bad Request",
	// 			Message: err.Error(),
	// 		})
	// 		return
	// 	}
	// }

	res, err := handler.commentService.FindByPhotoId(c, comment.PhotoID, uint(commentId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.ResponseErrorGeneral{
			Status:  "Data Not Found",
			Message: fmt.Sprintf("Comment Data with id: %d not found\n", commentId),
		})
		return
	}

	c.JSON(http.StatusOK, res)
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
// @Security token
// @securityDefinitions.apikey token
// @Param requestCreate body model.RequestComment true "Create Comment photo"
// @Success 201 {object} model.Comment
// @Failure 400 {object} model.ResponseErrorGeneral
// @Router /comment [post]
func (handler *commentHandler) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestComment{}

	// bind messange, photo_id
	var err error
	if contentType == appJson {
		err = c.ShouldBindJSON(&comment)
	} else {
		err = c.ShouldBind(&comment)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(comment); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range errors {
				switch e.Field() {
				case "Message":
					errMsg = "Invalid title."
				case "Caption":
					errMsg = "Invalid caption."
				case "PhotoID":
					errMsg = "Invalid photo_id."
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
	// 	if err := c.ShouldBindJSON(&comment); err != nil {
	// 		if err := handler.validate.Struct(comment); err != nil {
	// 			if errors, ok := err.(validator.ValidationErrors); ok {
	// 				var errMsg string
	// 				for _, e := range errors {
	// 					switch e.Field() {
	// 					case "Message":
	// 						errMsg = "Invalid title."
	// 					case "Caption":
	// 						errMsg = "Invalid caption."
	// 					case "PhotoID":
	// 						errMsg = "Invalid photo_id."
	// 					}
	// 				}
	// 				c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 					Status:  "Bad Request json",
	// 					Message: errMsg,
	// 				})
	// 				return
	// 			}
	// 		}
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&comment); err != nil {
	// 		if err := handler.validate.Struct(comment); err != nil {
	// 			if errors, ok := err.(validator.ValidationErrors); ok {
	// 				var errMsg string
	// 				for _, e := range errors {
	// 					switch e.Field() {
	// 					case "Message":
	// 						errMsg = "Invalid title."
	// 					case "Caption":
	// 						errMsg = "Invalid caption."
	// 					case "PhotoID":
	// 						errMsg = "Invalid photo_id."
	// 					}
	// 				}
	// 				c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 					Status:  "Bad Request form",
	// 					Message: errMsg,
	// 				})
	// 				return
	// 			}
	// 		}
	// 	}
	// }

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
// @Security token
// @securityDefinitions.apikey token
// @Param commentId path int true "Comment ID"
// @Param requestUpdate body model.RequestComment true "Update Comment by photo id user"
// @Success 200 {object} model.Comment
// @Failure 400 {object} model.ResponseErrorGeneral
// @Router /comment/{commentId} [put]
func (handler *commentHandler) UpdateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestComment{}

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
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(comment); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range errors {
				switch e.Field() {
				case "Message":
					errMsg = "Invalid title."
				case "Caption":
					errMsg = "Invalid caption."
				case "PhotoID":
					errMsg = "Invalid photo_id."
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
	// 	if err := c.ShouldBindJSON(&comment); err != nil {
	// 		if err := handler.validate.Struct(comment); err != nil {
	// 			if errors, ok := err.(validator.ValidationErrors); ok {
	// 				var errMsg string
	// 				for _, e := range errors {
	// 					switch e.Field() {
	// 					case "Message":
	// 						errMsg = "Invalid title."
	// 					case "Caption":
	// 						errMsg = "Invalid caption."
	// 					case "PhotoID":
	// 						errMsg = "Invalid photo_id."
	// 					}
	// 				}
	// 				c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 					Status:  "Bad Request json",
	// 					Message: errMsg,
	// 				})
	// 				return
	// 			}
	// 		}
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&comment); err != nil {
	// 		if err := handler.validate.Struct(comment); err != nil {
	// 			if errors, ok := err.(validator.ValidationErrors); ok {
	// 				var errMsg string
	// 				for _, e := range errors {
	// 					switch e.Field() {
	// 					case "Message":
	// 						errMsg = "Invalid title."
	// 					case "Caption":
	// 						errMsg = "Invalid caption."
	// 					case "PhotoID":
	// 						errMsg = "Invalid photo_id."
	// 					}
	// 				}
	// 				c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 					Status:  "Bad Request form",
	// 					Message: errMsg,
	// 				})
	// 				return
	// 			}
	// 		}
	// 	}
	// }

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
// @Security token
// @securityDefinitions.apikey token
// @Param commentId path int true "Comment ID"
// @Success 200 {object} model.ResponseDeleted
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 500 {object} model.ResponseErrorGeneral
// @Router /comment/{commentId} [delete]
func (handler *commentHandler) DeleteComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestDeleteComment{}

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
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(comment); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range errors {
				switch e.Field() {
				case "Message":
					errMsg = "Invalid title."
				case "Caption":
					errMsg = "Invalid caption."
				case "PhotoID":
					errMsg = "Invalid photo_id."
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
	// 	if err := c.ShouldBindJSON(&comment); err != nil {
	// 		if err := handler.validate.Struct(comment); err != nil {
	// 			if errors, ok := err.(validator.ValidationErrors); ok {
	// 				var errMsg string
	// 				for _, e := range errors {
	// 					switch e.Field() {
	// 					case "PhotoID":
	// 						errMsg = "Invalid photo_id."
	// 					}
	// 				}
	// 				c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 					Status:  "Bad Request json",
	// 					Message: errMsg,
	// 				})
	// 				return
	// 			}
	// 		}
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&comment); err != nil {
	// 		if err := handler.validate.Struct(comment); err != nil {
	// 			if errors, ok := err.(validator.ValidationErrors); ok {
	// 				var errMsg string
	// 				for _, e := range errors {
	// 					switch e.Field() {
	// 					case "PhotoID":
	// 						errMsg = "Invalid photo_id."
	// 					}
	// 				}
	// 				c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
	// 					Status:  "Bad Request form",
	// 					Message: errMsg,
	// 				})
	// 				return
	// 			}
	// 		}
	// 	}
	// }

	if _, err := handler.commentService.Delete(c, uint(commentId), userID, comment.PhotoID); err != nil {
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

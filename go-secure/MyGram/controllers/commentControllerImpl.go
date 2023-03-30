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

func NewCommentHandler(commentService services.CommentService, validator_ validator.Validate) *commentHandler {
	return &commentHandler{
		commentService: commentService,
		validate:       &validator_,
	}
}

// Get all comment by photo id
// all user can access to show all comment by photo id
func (handler *commentHandler) GetAllComment(c *gin.Context) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.Comment{}

	// bind photo_id
	if contentType == appJson {
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&comment); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	res, err := handler.commentService.FindAllByPhotoId(c, comment.PhotoID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": fmt.Sprintf("Error retrieving all comment user: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Get comment by photo id
// all user can access read comment by photo id
func (handler *commentHandler) GetComment(c *gin.Context) {
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.Comment{}

	// Get param commentId
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil || uint(commentId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid comment ID",
		})
		return
	}

	// bind photo_id
	if contentType == appJson {
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&comment); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	res, err := handler.commentService.FindByPhotoId(c, comment.PhotoID, uint(commentId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": fmt.Sprintf("Comment Data with id: %d not found\n", commentId),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Create comment by photo id and user id
// all user can create comment by photo id and user id 
func (handler *commentHandler) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestComment{}

	// bind messange, photo_id
	if contentType == appJson {
		if err := c.ShouldBindJSON(&comment); err != nil {
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
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request json",
					"message": errMsg,
				})
				return
			}
		}
	} else {
		if err := c.ShouldBind(&comment); err != nil {
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
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request form",
					"message": errMsg,
				})
				return
			}
		}
	}

	if res, err := handler.commentService.Create(c, comment, userID, comment.PhotoID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// Update comment by photo id and user id
// User can update comment by photo id and user id
func (handler *commentHandler) UpdateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestComment{}

	// Get param commentId
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid comment ID",
		})
		return
	}

	// bind message, photo_id
	if contentType == appJson {
		if err := c.ShouldBindJSON(&comment); err != nil {
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
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request json",
					"message": errMsg,
				})
				return
			}
		}
	} else {
		if err := c.ShouldBind(&comment); err != nil {
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
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request form",
					"message": errMsg,
				})
				return
			}
		}
	}

	if res, err := handler.commentService.Update(c, comment, uint(commentId), userID, comment.PhotoID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Delete comment by photo id and user id
// User can delete comment by photo id and user id
func (handler *commentHandler) DeleteComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	comment := model.RequestDeleteComment{}

	// Get param commentId
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil || uint(commentId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid comment ID",
		})
		return
	}

	// bind photo_id
	if contentType == appJson {
		if err := c.ShouldBindJSON(&comment); err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				var errMsg string
				for _, e := range errors {
					switch e.Field() {
					case "PhotoID":
						errMsg = "Invalid photo_id."
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
		if err := c.ShouldBind(&comment); err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				var errMsg string
				for _, e := range errors {
					switch e.Field() {
					case "PhotoID":
						errMsg = "Invalid photo_id."
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

	if _, err := handler.commentService.Delete(c, uint(commentId), userID, comment.PhotoID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete comment",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data Comment deleted successfully",
		})
	}
}

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

type photoHandler struct {
	photoService services.PhotoService
	validate     *validator.Validate
}

func NewPhotoHandler(photoService services.PhotoService, validator_ validator.Validate) *photoHandler {
	return &photoHandler{
		photoService: photoService,
		validate:     &validator_,
	}
}

// Get all photo by user id
// User can access to show all photo by user id
func (handler *photoHandler) GetAllPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	res, err := handler.photoService.FindAllByUserId(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": fmt.Sprintf("Error retrieving all photo user: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Get photo by user id
// User can access to show photo by user id
func (handler *photoHandler) GetPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get param photo_id
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil || uint(photoId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid photo ID",
		})
		return
	}

	res, err := handler.photoService.FindByUserId(c, userID, uint(photoId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": fmt.Sprintf("Photo Data with id: %d not found\n", photoId),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create photo by user id
// User can create photo by user id
func (handler *photoHandler) CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	photo := model.Photo{}

	// bind title, caption (optional), photo_url
	if contentType == appJson {
		if err := c.ShouldBindJSON(&photo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&photo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	// validate input json or form
	if err := handler.validate.Struct(photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if res, err := handler.photoService.Create(c, photo, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// Update photo by user id
// User can update photo by user id
func (handler *photoHandler) UpdatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	photo := model.Photo{}

	// Get Param PhotoID
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid photo ID",
		})
		return
	}

	// bind title, caption (optional), photo_url
	if contentType == appJson {
		if err := c.ShouldBindJSON(&photo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&photo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	// validate input json or form
	if err := handler.validate.Struct(photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if res, err := handler.photoService.Update(c, photo, uint(photoId), userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Delete photo by user id
// User can delete photo by user id
func (handler *photoHandler) DeletePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get Param PhotoId
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil || uint(photoId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid photo ID",
		})
		return
	}

	if _, err := handler.photoService.Delete(c, uint(photoId), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete photo",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data Photo deleted successfully",
		})
	}
}

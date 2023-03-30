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

type socialMediaHandler struct {
	socialMediaService services.SocialMediaService
	validate           *validator.Validate
}

func NewSocialMediaHandler(socialMediaService services.SocialMediaService, validator_ validator.Validate) *socialMediaHandler {
	return &socialMediaHandler{
		socialMediaService: socialMediaService,
		validate:           &validator_,
	}
}

// Get all social media by user id
// User can access to show all social media by user id
func (handler *socialMediaHandler) GetAllSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	res, err := handler.socialMediaService.FindAllByUserId(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": fmt.Sprintf("Error retrieving all social media user: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Get social media by user id
// User can access to show social media by user id
func (handler *socialMediaHandler) GetSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get Param socialMediaId
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil || uint(socialMediaId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid social media ID",
		})
		return
	}

	res, err := handler.socialMediaService.FindByUserId(c, userID, uint(socialMediaId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": fmt.Sprintf("Social Media Data with id: %d not found\n", socialMediaId),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create social media by user id
// User can create social media by user id
func (handler *socialMediaHandler) CreateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	socialMedia := model.SocialMedia{}

	// bind name, social_media_url
	if contentType == appJson {
		if err := c.ShouldBindJSON(&socialMedia); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&socialMedia); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	// validate input json or form
	if err := handler.validate.Struct(socialMedia); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if res, err := handler.socialMediaService.Create(c, socialMedia, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// Update social media by user id
// User can update social media by user id
func (handler *socialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	socialMedia := model.SocialMedia{}

	// Get Param socialMediaID
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid social media ID",
		})
		return
	}

	// bind name, social_media_url
	if contentType == appJson {
		if err := c.ShouldBindJSON(&socialMedia); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&socialMedia); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	// validate input json or form
	if err := handler.validate.Struct(socialMedia); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if res, err := handler.socialMediaService.Update(c, socialMedia, uint(socialMediaId), userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Delete social media by user id
// User can delete social media by user id
func (handler *socialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get Param socialMediaId
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil || uint(socialMediaId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid social media ID",
		})
		return
	}

	if _, err := handler.socialMediaService.Delete(c, uint(socialMediaId), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete social media",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data Social Media deleted successfully",
		})
	}
}

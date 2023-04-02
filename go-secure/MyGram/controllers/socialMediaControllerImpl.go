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
// Get all social media godoc
// @Summary Get all social media user
// @Description Get all social media user
// @Tags Social Media
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Success 200 {array} model.SocialMedia
// @Failure 500 {object} model.ResponseErrorGeneral
// @Router /socialmedia [get]
func (handler *socialMediaHandler) GetAllSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	res, err := handler.socialMediaService.FindAllByUserId(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseErrorGeneral{
			Status:  "Internal Server Error",
			Message: fmt.Sprintf("Error retrieving all social media user: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Get social media by user id
// User can access to show social media by user id
// Get social media godoc
// @Summary Get social media user
// @Description Get social media user
// @Tags Social Media
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param socialMediaId path int true "Social Media ID"
// @Success 200 {object} model.SocialMedia
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 404 {object} model.ResponseErrorGeneral
// @Router /socialmedia/{socialMediaId} [get]
func (handler *socialMediaHandler) GetSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get Param socialMediaId
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil || uint(socialMediaId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: "invalid social media ID",
		})
		return
	}

	res, err := handler.socialMediaService.FindByUserId(c, userID, uint(socialMediaId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.ResponseErrorGeneral{
			Status:  "Data Not Found",
			Message: fmt.Sprintf("Social Media Data with id: %d not found\n", socialMediaId),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create social media by user id
// User can create social media by user id
// Create social media godoc
// @Summary Create social media user
// @Description Create social media user
// @Tags Social Media
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param requestCreate body model.RequestSocialMedia true "Create Social Media user"
// @Success 201 {object} model.SocialMedia
// @Failure 400 {object} model.ResponseErrorGeneral
// @Router /socialmedia [post]
func (handler *socialMediaHandler) CreateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	socialMedia := model.RequestSocialMedia{}

	// bind name, social_media_url
	var err error
	if contentType == appJson {
		err = c.ShouldBindJSON(&socialMedia)
	} else {
		err = c.ShouldBind(&socialMedia)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(socialMedia); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range errors {
				switch e.Field() {
				case "Name":
					errMsg = "Invalid name."
				case "SocialMediaURL":
					errMsg = "Invalid url."
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
	// 	if err := c.ShouldBindJSON(&socialMedia); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Name":
	// 					errMsg = "Invalid name."
	// 				case "SocialMediaURL":
	// 					errMsg = "Invalid url."
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
	// 	if err := c.ShouldBind(&socialMedia); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Name":
	// 					errMsg = "Invalid name."
	// 				case "SocialMediaURL":
	// 					errMsg = "Invalid url."
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

	if res, err := handler.socialMediaService.Create(c, socialMedia, userID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// Update social media by user id
// User can update social media by user id
// Update social media godoc
// @Summary Update social media user
// @Description Update social media user
// @Tags Social Media
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param socialMediaId path int true "Social Media ID"
// @Param requestUpdate body model.RequestSocialMedia true "Update Social Media user"
// @Success 200 {object} model.SocialMedia
// @Failure 400 {object} model.ResponseErrorGeneral
// @Router /socialmedia/{socialMediaId} [put]
func (handler *socialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	socialMedia := model.RequestSocialMedia{}

	// Get Param socialMediaID
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request form",
			Message: "invalid social media ID",
		})
		return
	}

	// bind name, social_media_url
	if contentType == appJson {
		err = c.ShouldBindJSON(&socialMedia)
	} else {
		err = c.ShouldBind(&socialMedia)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(socialMedia); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range errors {
				switch e.Field() {
				case "Name":
					errMsg = "Invalid name."
				case "SocialMediaURL":
					errMsg = "Invalid url."
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
	// 	if err := c.ShouldBindJSON(&socialMedia); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Name":
	// 					errMsg = "Invalid name."
	// 				case "SocialMediaURL":
	// 					errMsg = "Invalid url."
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
	// 	if err := c.ShouldBind(&socialMedia); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Name":
	// 					errMsg = "Invalid name."
	// 				case "SocialMediaURL":
	// 					errMsg = "Invalid url."
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

	if res, err := handler.socialMediaService.Update(c, socialMedia, uint(socialMediaId), userID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Delete social media by user id
// User can delete social media by user id
// Delete social media godoc
// @Summary Delete social media user
// @Description Delete social media user
// @Tags Social Media
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param socialMediaId path int true "Social Media ID"
// @Success 200 {object} model.ResponseDeleted
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 500 {object} model.ResponseErrorGeneral
// @Router /socialmedia/{socialMediaId} [delete]
func (handler *socialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get Param socialMediaId
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil || uint(socialMediaId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: "invalid social media ID",
		})
		return
	}

	if _, err := handler.socialMediaService.Delete(c, uint(socialMediaId), userID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseErrorGeneral{
			Status:  "Internal Server Error",
			Message: "Failed to delete social media",
		})
		return
	} else {
		c.JSON(http.StatusOK, model.ResponseDeleted{
			Message: "Data Social Media deleted successfully",
		})
	}
}

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
// Get all photo godoc
// @Summary Get all photo user
// @Description Get all photo user
// @Tags Photo
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Success 200 {array} model.Photo
// @Failure 500 {object} model.ResponseErrorGeneral
// @Router /photo [get]
func (handler *photoHandler) GetAllPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	res, err := handler.photoService.FindAllByUserId(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.ResponseErrorGeneral{
			Status:  "Internal Server Error",
			Message: fmt.Sprintf("Error retrieving all photo user: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Get photo by user id
// User can access to show photo by user id
// Get photo godoc
// @Summary Get photo user
// @Description Get photo user
// @Tags Photo
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param photoId path int true "Photo ID"
// @Success 200 {object} model.Photo
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 404 {object} model.ResponseErrorGeneral
// @Router /photo/{photoId} [get]
func (handler *photoHandler) GetPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get param photo_id
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil || uint(photoId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: "invalid photo ID",
		})

		return
	}

	res, err := handler.photoService.FindByUserId(c, userID, uint(photoId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.ResponseErrorGeneral{
			Status:  "Data Not Found",
			Message: fmt.Sprintf("Photo Data with id: %d not found\n", photoId),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create photo by user id
// User can create photo by user id
// Create photo godoc
// @Summary Create photo user
// @Description Create photo user
// @Tags Photo
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param requestCreate body model.RequestPhoto true "Create Photo user"
// @Success 201 {object} model.Photo
// @Failure 400 {object} model.ResponseErrorGeneral
// @Router /photo [post]
func (handler *photoHandler) CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	photo := model.RequestPhoto{}

	// bind title, caption (optional), photo_url
	var err error
	if contentType == appJson {
		err = c.ShouldBindJSON(&photo)
	} else {
		err = c.ShouldBind(&photo)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(photo); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range errors {
				switch e.Field() {
				case "Title":
					errMsg = "Invalid title."
				case "Caption":
					errMsg = "Invalid caption."
				case "PhotoURL":
					errMsg = "Invalid photo url."
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
	// 	if err := c.ShouldBindJSON(&photo); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Title":
	// 					errMsg = "Invalid title."
	// 				case "Caption":
	// 					errMsg = "Invalid caption."
	// 				case "PhotoURL":
	// 					errMsg = "Invalid photo url."
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
	// 	if err := c.ShouldBind(&photo); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Title":
	// 					errMsg = "Invalid title."
	// 				case "Caption":
	// 					errMsg = "Invalid caption."
	// 				case "PhotoURL":
	// 					errMsg = "Invalid photo url."
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

	if res, err := handler.photoService.Create(c, photo, userID); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// Update photo by user id
// User can update photo by user id
// Update photo godoc
// @Summary Update photo user
// @Description Update photo user
// @Tags Photo
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Produce x-www-form-urlencoded
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param photoId path int true "Photo ID"
// @Param requestUpdate body model.RequestPhoto true "Update Photo user"
// @Success 200 {object} model.Photo
// @Failure 400 {object} model.ResponseErrorGeneral
// @Router /photo/{photoId} [put]
func (handler *photoHandler) UpdatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	photo := model.RequestPhoto{}

	// Get Param PhotoID
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: "invalid photo ID",
		})
		return
	}

	// bind title, caption (optional), photo_url
	if contentType == appJson {
		err = c.ShouldBindJSON(&photo)
	} else {
		err = c.ShouldBind(&photo)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := handler.validate.Struct(photo); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, e := range errors {
				switch e.Field() {
				case "Title":
					errMsg = "Invalid title."
				case "Caption":
					errMsg = "Invalid caption."
				case "PhotoURL":
					errMsg = "Invalid photo url."
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
	// 	if err := c.ShouldBindJSON(&photo); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Title":
	// 					errMsg = "Invalid title."
	// 				case "Caption":
	// 					errMsg = "Invalid caption."
	// 				case "PhotoURL":
	// 					errMsg = "Invalid photo url."
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
	// 	if err := c.ShouldBind(&photo); err != nil {
	// 		if errors, ok := err.(validator.ValidationErrors); ok {
	// 			var errMsg string
	// 			for _, e := range errors {
	// 				switch e.Field() {
	// 				case "Title":
	// 					errMsg = "Invalid title."
	// 				case "Caption":
	// 					errMsg = "Invalid caption."
	// 				case "PhotoURL":
	// 					errMsg = "Invalid photo url."
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

	if res, err := handler.photoService.Update(c, photo, uint(photoId), userID); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Delete photo by user id
// User can delete photo by user id
// Delete photo godoc
// @Summary Delete photo user
// @Description Delete photo user
// @Tags Photo
// @Accept json
// @Produce json
// @Security JWT
// @securityDefinitions.apikey JWT
// @Param photoId path int true "Photo ID"
// @Success 200 {object} model.ResponseDeleted
// @Failure 400 {object} model.ResponseErrorGeneral
// @Failure 500 {object} model.ResponseErrorGeneral
// @Router /photo/{photoId} [delete]
func (handler *photoHandler) DeletePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get Param PhotoId
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil || uint(photoId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
			Status:  "Bad Request",
			Message: "invalid photo ID",
		})
		return
	}

	if _, err := handler.photoService.Delete(c, uint(photoId), userID); err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseErrorGeneral{
			Status:  "Internal Server Error",
			Message: "Failed to delete photo",
		})
		return
	} else {
		c.JSON(http.StatusOK, model.ResponseDeleted{
			Message: "Data Photo deleted successfully",
		})
	}
}

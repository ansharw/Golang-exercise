package controllers

import (
	"challenges-two/database"
	"challenges-two/helpers"
	"challenges-two/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Admin
// GET, GET ALL, UPDATE, DELETE, POST
func GetAllProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	product := models.Product{}
	products := []models.Product{}
	userID := uint(userData["id"].(float64))

	product.UserID = userID

	if userID == 1 {
		result := db.Order("id DESC").Find(&products)
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(404, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintln("There is no product"),
			})
			return
		}

		c.JSON(http.StatusOK, products)
	}

	// result := db.Debug().Where("user_id = ?", userID).Find(&product)
	result := db.Order("id DESC").Where("user_id = ?", userID).Find(&products)

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"error":   "Data Not Found",
			"message": fmt.Sprintln("There is no product"),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if userID == 1 {
		err := db.Where("id = ?", productId).First(&product).Error
		if err != nil || productId == 0 {
			c.AbortWithStatusJSON(404, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintf("product with id: %d not found\n", productId),
			})
			return
		}
	}

	// err := db.Debug().Where("id = ? AND user_id = ?", productId, userID).First(&product).Error
	err := db.Where("id = ? AND user_id = ?", productId, userID).First(&product).Error

	if err != nil || productId == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"error":   "Data Not Found",
			"message": fmt.Sprintf("product with id: %d not found\n", productId),
		})
		return
	}

	product.UserID = userID
	product.ID = uint(productId)

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID

	// err := db.Debug().Create(&product).Error
	err := db.Create(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// only admin has access
func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	// untuk user selain admin di tolak
	// ONLY ADMIN HAS ACCESS
	if userID != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Access denied",
			"message": "You hasn't access to use this feature",
		})
		return
	}

	if contentType == appJson {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID
	product.ID = uint(productId)

	err := db.Model(&product).Where("id = ?", productId).Updates(models.Product{Title: product.Title, Description: product.Description}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

// only admin has access
func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	// untuk user selain admin di tolak
	// ONLY ADMIN HAS ACCESS
	if userID != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Access denied",
			"message": "You hasn't access to use this feature",
		})
		return
	}

	if contentType == appJson {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	err_ := db.Model(&product).Where("id = ?", productId).First(&product).Error
	if err_ != nil || productId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data not found",
			"message": "Product doesn't exist",
		})
		return
	}

	product.UserID = userID
	product.ID = uint(productId)

	err := db.Model(&product).Where("id = ?", productId).Delete(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

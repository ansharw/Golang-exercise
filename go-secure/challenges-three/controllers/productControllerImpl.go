package controllers

import (
	"challenges-three/database"
	"challenges-three/helpers"
	"challenges-three/models"
	"challenges-three/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

// Admin
// GET, GET ALL, UPDATE, DELETE, POST
// User
// GET, GET ALL, POST
func (handler *productHandler) GetAllProducts(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if userID == 1 {
		res, err := handler.productService.FindAll(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": fmt.Sprintf("Error retrieving products: %s", err.Error()),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := handler.productService.FindAllByUserId(c, userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": fmt.Sprintf("Error retrieving products: %s", err.Error()),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}


func GetProduct(c *gin.Context) {
	db := database.GetConnection()
	userData := c.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid product ID",
		})
		return
	}
	userID := uint(userData["id"].(float64))

	if userID == 1 {
		err := db.Where("id = ?", productId).First(&product).Error
		if err != nil || productId == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintf("Product with id: %d not found\n", productId),
			})
			return
		}
	}

	// err := db.Debug().Where("id = ? AND user_id = ?", productId, userID).First(&product).Error
	err_ := db.Where("id = ? AND user_id = ?", productId, userID).First(&product).Error

	if err_ != nil || productId == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": fmt.Sprintf("Product with id: %d not found\n", productId),
		})
		return
	}

	product.UserID = userID
	product.ID = uint(productId)

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	db := database.GetConnection()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	product := models.Product{}
	userID := uint(userData["id"].(float64))

	// set initial userid for post method
	product.UserID = userID

	if contentType == appJson {
		if err := c.ShouldBindJSON(&product); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&product); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	// err := db.Debug().Create(&product).Error
	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// only admin to access this feature
func UpdateProduct(c *gin.Context) {
	db := database.GetConnection()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	// untuk user selain admin di tolak
	// ONLY ADMIN HAS ACCESS
	userID := uint(userData["id"].(float64))
	if userID != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Access denied",
			"message": "You do not have permission to access this feature",
		})
		return
	}

	product := models.Product{}
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid product ID",
		})
		return
	}

	if contentType == appJson {
		if err := c.ShouldBindJSON(&product); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&product); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	product.ID = uint(productId)

	err_ := db.Model(&product).Where("id = ?", productId).Updates(models.Product{Title: product.Title, Description: product.Description}).Error

	if err_ != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

// only admin to access this feature
func DeleteProduct(c *gin.Context) {
	db := database.GetConnection()
	userData := c.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}

	// untuk user selain admin di tolak
	// ONLY ADMIN HAS ACCESS
	userID := uint(userData["id"].(float64))
	if userID != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Access denied",
			"message": "You do not have permission to access this feature",
		})
		return
	}

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid product ID",
		})
		return
	}

	err = db.Model(&product).Where("id = ?", productId).First(&product).Error
	if err != nil || productId == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": "Product doesn't exist",
		})
		return
	}

	err = db.Model(&product).Where("id = ?", productId).Delete(&product).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

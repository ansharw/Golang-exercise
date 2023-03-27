package controllers

import (
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
	userID := userData["id"].(uint)

	if userID == 1 {
		res, err := handler.productService.FindAll(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": fmt.Sprintf("Error retrieving products: %s", err.Error()),
			})
			return
		}
		fmt.Println(res)
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

func (handler *productHandler) GetProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(uint)

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil || uint(productId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid product ID",
		})
		return
	}

	if userID == 1 {
		res, err := handler.productService.FindById(c, uint(productId))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintf("Product with id: %d not found\n", productId),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := handler.productService.FindAllByUserId(c, userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintf("Product with id: %d not found\n", productId),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func (handler *productHandler) CreateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(uint)
	contentType := helpers.GetContentType(c)
	product := models.Product{}

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

	if res, err := handler.productService.Create(c, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// only admin to access this feature
func (handler *productHandler) UpdateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(uint)
	contentType := helpers.GetContentType(c)
	product := models.Product{}

	// untuk user selain admin di tolak
	// ONLY ADMIN HAS ACCESS
	if userID != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Access denied",
			"message": "You do not have permission to access this feature",
		})
		return
	}

	// error ProductID
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

	if res, err := handler.productService.Update(c, uint(productId)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// only admin to access this feature
func (handler *productHandler) DeleteProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(uint)

	// untuk user selain admin di tolak
	// ONLY ADMIN HAS ACCESS
	if userID != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Access denied",
			"message": "You do not have permission to access this feature",
		})
		return
	}

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil || uint(productId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid product ID",
		})
		return
	}

	if _, err := handler.productService.Delete(c, uint(productId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete product",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product deleted successfully",
		})
	}
}

package middlewares

import (
	"challenges-two/database"
	"challenges-two/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// func ProductAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := database.GetDB()
// 		productID, err := strconv.Atoi(c.Param("productId"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error":   "Bad Request",
// 				"message": "invalid parameter",
// 			})
// 			return
// 		}

// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		userID := uint(userData["id"].(float64))
// 		product := models.Product{}

// 		err = db.Select("user_id").First(&product, uint(productID)).Error

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Data not found",
// 				"message": "Data doesn't exist",
// 			})
// 			return
// 		}

// 		if product.UserID != userID {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this data",
// 			})
// 			return
// 		}

// 		c.Next()
// 	}
// }

// without parameter
func ProductAuthorizations() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		products := []models.Product{}

		if userID == 1 {
			result := db.Order("id desc").Find(&products)
			if result.RowsAffected == 0 {
				c.AbortWithStatusJSON(404, gin.H{
					"error":   "Data Not Found",
					"message": fmt.Sprintln("There is no product"),
				})
				return
			}
		}

		result := db.Where("user_id = ?", userID).Order("id desc").Find(&products)
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(404, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintln("There is no product"),
			})
			return
		}

		if products[0].UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

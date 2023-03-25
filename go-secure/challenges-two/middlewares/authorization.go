package middlewares

import (
	"challenges-two/database"
	"challenges-two/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ProductAuthorizations() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		product := models.Product{}

		if userID == 1 {
			result := db.Order("id desc").First(&product)
			if result.RowsAffected == 0 {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": fmt.Sprintln("There is no data"),
				})
				return
			}
		}

		result := db.Where("user_id = ?", userID).Order("id desc").First(&product)
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintln("There is no data"),
			})
			return
		}

		if product.UserID != userID && userID != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

package middlewares

import (
	"MyGram/database"
	"MyGram/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SocialMediaAuthorizations() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetConnection()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		socialMedia := model.SocialMedia{}

		result := db.Where("user_id = ?", userID).Order("id desc").Take(&socialMedia)
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintln("There is no data social media"),
			})
			return
		}

		// for protect unregistered user to login
		if socialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func CommentAuthorizations() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetConnection()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := model.Comment{}

		result := db.Where("user_id = ?", userID).Order("id desc").Take(&comment)
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintln("There is no data comment"),
			})
			return
		}

		// for protect unregistered user to login
		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func PhotoAuthorizations() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetConnection()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photo := model.Photo{}

		result := db.Where("user_id = ?", userID).Order("id desc").Take(&photo)
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintln("There is no data photo"),
			})
			return
		}

		// for protect unregistered user to login
		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
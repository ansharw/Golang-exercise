package middlewares

import (
	"MyGram/database"
	"MyGram/model"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SocialMediaAuthorizations(mdl interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetConnection()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		socialMediaID, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
				Status:  "Bad Request",
				Message: "invalid social media id",
			})
			return
		}

		valueOf := reflect.ValueOf(mdl)
		if valueOf.Kind() != reflect.Ptr {
			log.Println("model must be a pointer to a struct")
		}
		socialMedia__ := reflect.New(valueOf.Elem().Type()).Interface()

		err = db.Where("id = ?", socialMediaID).First(&socialMedia__).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, model.ResponseErrorGeneral{
				Status:  "Not Found",
				Message: "Social media not found",
			})
			return
		}

		// for protect unregistered user to login
		socialMediaValue := reflect.ValueOf(socialMedia__).Elem()

		socialMediaUserIDField := socialMediaValue.FieldByName("UserID")
		if !socialMediaUserIDField.IsValid() {
			log.Println("struct must have a field named UserID")
		}
		socialMediaUserID := uint(socialMediaUserIDField.Uint())

		if socialMediaUserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ResponseErrorGeneral{
				Status:  "Unauthorized",
				Message: "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func CommentAuthorizations(mdl interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetConnection()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		commentID, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
				Status:  "Bad Request",
				Message: "invalid comment id",
			})
			return
		}

		valueOf := reflect.ValueOf(mdl)
		if valueOf.Kind() != reflect.Ptr {
			log.Println("model must be a pointer to a struct")
		}
		comment__ := reflect.New(valueOf.Elem().Type()).Interface()

		err = db.Where("id = ?", commentID).First(&comment__).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, model.ResponseErrorGeneral{
				Status:  "Not Found",
				Message: "Comment not found",
			})
			return
		}

		// result := db.Where("user_id = ?", userID).Order("id desc").Take(&comment)
		// if result.RowsAffected == 0 {
		// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		// 		"error":   "Data Not Found",
		// 		"message": fmt.Sprintln("There is no data comment"),
		// 	})
		// 	return
		// }

		// for protect unregistered user to login
		commentValue := reflect.ValueOf(comment__).Elem()

		commentUserIDField := commentValue.FieldByName("UserID")
		if !commentUserIDField.IsValid() {
			log.Println("struct must have a field named UserID")
		}
		commentUserID := uint(commentUserIDField.Uint())

		if commentUserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ResponseErrorGeneral{
				Status:  "Unauthorized",
				Message: "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func PhotoAuthorizations(mdl interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetConnection()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		photoID, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseErrorGeneral{
				Status:  "Bad Request",
				Message: "invalid photo id",
			})
			return
		}

		valueOf := reflect.ValueOf(mdl)
		if valueOf.Kind() != reflect.Ptr {
			log.Println("model must be a pointer to a struct")
		}
		photo__ := reflect.New(valueOf.Elem().Type()).Interface()

		err = db.Where("id = ?", photoID).First(&photo__).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, model.ResponseErrorGeneral{
				Status:  "Not Found",
				Message: "Photo not found",
			})
			return
		}

		// for protect unregistered user to login
		photoValue := reflect.ValueOf(photo__).Elem()

		photoUserIDField := photoValue.FieldByName("UserID")
		if !photoUserIDField.IsValid() {
			panic("struct must have a field named UserID")
		}
		photoUserID := uint(photoUserIDField.Uint())

		if photoUserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ResponseErrorGeneral{
				Status:  "Unauthorized",
				Message: "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

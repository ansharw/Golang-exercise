package controllers

import (
	"challenges-three/database"
	"challenges-three/helpers"
	"challenges-three/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetConnection()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJson {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	// if email already exist
	var existingUser models.User
	if err := db.Where("email = ?", User.Email).First(&existingUser).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Email Already Exists",
			"message": "The email address you entered already exists",
		})
		return
	}

	// err := db.Debug().Create(&User).Error
	if err := db.Create(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.Fullname,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetConnection()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	Password := ""

	if contentType == appJson {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	Password = User.Password

	// err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	err := db.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(Password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

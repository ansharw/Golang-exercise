package controllers

import "github.com/gin-gonic/gin"

type PhotoHandler interface {
	GetAllPhoto(c *gin.Context)
	GetPhoto(c *gin.Context)
	CreatePhoto(c *gin.Context)
	UpdatePhoto(c *gin.Context)
	DeletePhoto(c *gin.Context)
}

package controllers

import "github.com/gin-gonic/gin"

type CommentHandler interface {
	GetAllComment(c *gin.Context)
	GetComment(c *gin.Context)
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
}

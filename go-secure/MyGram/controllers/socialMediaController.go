package controllers

import "github.com/gin-gonic/gin"

type SocialMediaHandler interface {
	GetAllSocialMedia(c *gin.Context)
	GetSocialMedia(c *gin.Context)
	CreateSocialMedia(c *gin.Context)
	UpdateSocialMedia(c *gin.Context)
	DeleteSocialMedia(c *gin.Context)
}

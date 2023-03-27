package controllers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

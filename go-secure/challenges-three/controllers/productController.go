package controllers

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	GetAllProducts(c *gin.Context)
}

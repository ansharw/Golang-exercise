package controllers

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	GetAllProducts(c *gin.Context)
	GetProduct(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

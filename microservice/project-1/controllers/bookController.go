package controllers

import (
	"fmt"
	"net/http"
	"project-1/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Book
		var req model.RequestBook
		var res model.ResponseBook

		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"status":  "Bad Request",
				"message": fmt.Sprintln("Bad Request JSON format"),
			})
			return
		}

		book.NameBook = req.NameBook
		book.Author = req.Author

		if err := db.Create(&book).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		res = model.ResponseBook(book)

		c.JSON(http.StatusCreated, res)
	}
}

func GetAllBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book []model.Book

		result := db.Find(&book)

		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(404, gin.H{
				"status":  "Data Not Found",
				"message": fmt.Sprintln("There is no book"),
			})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

func GetBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("bookID")
		bid, _ := strconv.Atoi(bookID)
		var book model.Book
		var res model.ResponseBook

		err := db.First(&book, bid).Error

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{
				"status":  "Data Not Found",
				"message": fmt.Sprintf("book with id %v not found", bid),
			})
			return
		}
		res = model.ResponseBook(book)
		c.JSON(http.StatusOK, res)
	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("bookID")
		bid, _ := strconv.Atoi(bookID)
		var book model.Book

		err := db.Delete(&book, bid).Error

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{
				"status":  "Data Not Found",
				"message": fmt.Sprintf("book with id %v not found", bid),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book deleted successfully",
		})
	}
}

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("bookID")
		bid, _ := strconv.Atoi(bookID)
		var book model.Book
		var res model.ResponseBook

		err := db.First(&book, bid).Error

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{
				"status":  "Data Not Found",
				"message": fmt.Sprintf("book with id %v not found", bid),
			})
			return
		}

		var req model.RequestBook
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request payload",
			})
			return
		}

		if req.NameBook != "" {
			book.NameBook = req.NameBook
		}
		if req.Author != "" {
			book.Author = req.Author
		}

		err = db.Model(&book).Updates(book).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		res = model.ResponseBook(book)
		c.JSON(http.StatusOK, res)
	}
}

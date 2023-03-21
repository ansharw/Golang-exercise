package controllers

import (
	"fmt"
	"net/http"
	"showcase/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddBook godoc
// @Summary Add book details
// @Description Add book details
// @Tags books
// @Accept json
// @Produce json
// @Param models.RequestBook body models.RequestBook true "Add the book"
// @Success 200 {object} models.ResponseBook
// @Failure 400 {object} models.BadRequest
// @Failure 500 {string} string "Internal Server Error"
// @Router /books [post]
func AddBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		var req models.RequestBook
		var res models.ResponseBook

		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(400, models.BadRequest{
				Status:  "Bad Request",
				Message: fmt.Sprintln("Bad Request JSON format"),
			})
			return
		}

		book.NameBook = req.NameBook
		book.Author = req.Author

		if err := db.Create(&book).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		res = models.ResponseBook(book)

		c.JSON(http.StatusCreated, res)
	}
}

// GetAllBook godoc
// @Summary Get details
// @Description Get details of all book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} models.ResponseBook
// @Failure 404 {object} models.NotFound
// @Router /books [get]
func GetAllBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book []models.Book

		result := db.Find(&book)

		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(404, models.NotFound{
				Status:  "Data Not Found",
				Message: fmt.Sprintln("There is no book"),
			})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

// GetBook godoc
// @Summary Get details by id
// @Description Get details of book by id
// @Tags books
// @Accept json
// @Produce json
// @Param BookID path int true "ID of the book"
// @Success 200 {object} models.ResponseBook
// @Failure 404 {object} models.NotFound
// @Router /books/{bookID} [get]
func GetBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("bookID")
		bid, _ := strconv.Atoi(bookID)
		var book models.Book
		var res models.ResponseBook

		err := db.First(&book, bid).Error

		if err != nil {
			c.AbortWithStatusJSON(404, models.NotFound{
				Status:  "Data Not Found",
				Message: fmt.Sprintf("book with id %v not found", bid),
			})
			return
		}
		res = models.ResponseBook(book)
		c.JSON(http.StatusOK, res)
	}
}

// DeleteBook godoc
// @Summary Delete by id
// @Description Delete of book by id
// @Tags books
// @Accept json
// @Produce json
// @Param BookID path int true "ID of the book"
// @Success 200 {object} models.Deleted
// @Failure 404 {object} models.NotFound
// @Router /books/{bookID} [delete]
func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("bookID")
		bid, _ := strconv.Atoi(bookID)
		var book models.Book

		err := db.Delete(&book, bid).Error

		if err != nil {
			c.AbortWithStatusJSON(404, models.NotFound{
				Status:  "Data Not Found",
				Message: fmt.Sprintf("book with id %v not found", bid),
			})
			return
		}

		c.JSON(http.StatusOK, models.Deleted{
			Message: "Book deleted successfully",
		})
	}
}

// UpdateBook godoc
// @Summary Update of the book by id
// @Description Update of the book by id
// @Tags books
// @Accept json
// @Produce json
// @Param BookID path int true "ID of the book"
// @Param models.RequestBook body models.RequestBook true "Update the book"
// @Success 200 {object} models.ResponseBook
// @Failure 404 {object} models.NotFound
// @Failure 400 {object} models.BadRequest
// @Failure 500 {string} string "Internal Server Error"
// @Router /books/{bookID} [put]
func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("bookID")
		bid, _ := strconv.Atoi(bookID)
		var book models.Book
		var res models.ResponseBook

		err := db.First(&book, bid).Error

		if err != nil {
			c.AbortWithStatusJSON(404, models.NotFound{
				Status:  "Data Not Found",
				Message: fmt.Sprintf("book with id %v not found", bid),
			})
			return
		}

		var req models.RequestBook
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.BadRequest{
				Status:  "Bad Request",
				Message: fmt.Sprintln("Bad Request JSON format"),
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

		res = models.ResponseBook(book)
		c.JSON(http.StatusOK, res)
	}
}

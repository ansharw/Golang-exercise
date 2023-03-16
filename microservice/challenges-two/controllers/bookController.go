package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func AddBook(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	c.JSON(http.StatusCreated, "Created")
}

func UpdateBook(c *gin.Context) {
	bookID := c.Param("bookID")
	bookid, _ := strconv.Atoi(bookID)
	condition := false
	var updateBook Book

	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookid == book.BookID {
			condition = true
			BookDatas[i] = updateBook
			BookDatas[i].BookID = bookid
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookid),
		})
		return
	}

	c.JSON(http.StatusOK, "Updated")
}

func GetBookById(c *gin.Context) {
	bookID := c.Param("bookID")
	bookid, _ := strconv.Atoi(bookID)
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if bookid == book.BookID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with ID %v not found", bookid),
		})
		return
	}

	c.JSON(http.StatusOK, bookData)
}

func GetAllBook(c *gin.Context) {
	condition := len(BookDatas)

	if condition == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintln("there is no book"),
		})
		return
	}

	c.JSON(http.StatusOK, BookDatas)
}

func DeleteBook(c *gin.Context) {
	bookID := c.Param("bookID")
	bookid, _ := strconv.Atoi(bookID)
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookid == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with ID %v not found", bookid),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	c.JSON(http.StatusOK, "Deleted")
}

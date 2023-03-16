package controllers

import (
	"challenges-three/model"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		// fmt.Println(book)

		sqlStatement := `INSERT INTO book (title, author, description) VALUES ($1, $2, $3) RETURNING *;`
		err := db.QueryRow(sqlStatement, book.Title, book.Author, book.Desc).Scan(&book.Id, &book.Title, &book.Author, &book.Desc)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, "Created")
	}
}

func GetAllBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []model.Book

		sqlStatement := `SELECT * FROM book;`
		rows, err := db.Query(sqlStatement)

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			var book model.Book
			err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc)

			books = append(books, book)
		}

		if (len(books) == 0 || err != nil) {
			// OK but no content
			c.AbortWithStatusJSON(204, gin.H{
				"status": "Data Not Found",
				"message": fmt.Sprintln("There is no book"),
			})
			return
		}
		
		c.JSON(http.StatusOK, books)
	}
}

func GetBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("bookID")
		bid, _ := strconv.Atoi(bookID)
		var book model.Book

		sqlStatement := `SELECT id, title, author, description FROM book WHERE id = $1;`
		err := db.QueryRow(sqlStatement, bid).Scan(&book.Id, &book.Title, &book.Author, &book.Desc)

		if err != nil {
			// OK but no content
			c.AbortWithStatusJSON(204, gin.H{
				"status": "Data Not Found",
				"message": fmt.Sprintf("book with id %v not found", bid),
			})
			return
		}
		
		c.JSON(http.StatusOK, book)
	}
}

func UpdateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("bookID")
		bid, _ := strconv.Atoi(bookID)
		var book model.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		// fmt.Println(book)

		sqlStatement := `UPDATE book SET title = $2, author = $3, description = $4 WHERE id = $1;`
		_, err := db.Exec(sqlStatement, bid, &book.Title, &book.Author, &book.Desc)

		fmt.Println(err)
		if err != nil {
			// OK but no content
			c.AbortWithStatusJSON(204, gin.H{
				"status": "Data Not Found",
				"message": fmt.Sprintf("book with id %v not found", bid),
			})
			return
		}
		
		c.JSON(http.StatusCreated, "Updated")
	}
}

func DeleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("bookID")
		bid, _ := strconv.Atoi(bookID)

		sqlStatement := `DELETE FROM book WHERE book.id = $1`
		_, err := db.Exec(sqlStatement, bid)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "Data Not Found",
				"message": fmt.Sprintf("book with id %v not found", bid),
			})
			return
		}

		c.JSON(http.StatusCreated, "Deleted")
	}
}

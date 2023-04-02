package controllers

import (
	"net/http"
	"sql_api_implementation_2/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

func CreateBook(c *gin.Context) {
	// set objek book
	var newBook *Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
		INSERT INTO books(title, author, description)
		VALUES($1, $2, $3)
		Returning *
	`
	_, err := config.DB.Exec(sqlStatement, newBook.Title, newBook.Author, newBook.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data is not found",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"status": "created",
	})
}

func GetAllBook(c *gin.Context) {
	books := []Book{}

	sqlStatement := `
	SELECT * FROM books
	`

	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"books": books,
	})
}

func GetBookByID(c *gin.Context) {
	bookID := c.Param("bookID")
	id, _ := strconv.Atoi(bookID)

	book := Book{}

	sqlStatement := `
	SELECT * FROM books
	WHERE id = $1
	`

	rows, err := config.DB.Query(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data is not found",
		})
		return
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "status bad gateaway",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"book": book,
	})
}

func UpdateBookByID(c *gin.Context) {
	id := c.Param("bookID")
	book := Book{}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannon decode json into struct",
		})
		return
	}
	
	sqlStatement := `
	UPDATE books 
	SET title=$1, author=$2, description=$3
	WHERE id=$4;
	`

	rows, err := config.DB.Exec(sqlStatement, book.Title, book.Author, book.Description, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot updated data",
		})
		return
	}

	count, _ := rows.RowsAffected()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"updated data amount": count,
	})
}

func DeleteBookByID(c *gin.Context) {
	id := c.Param("bookID")
	sqlStatement := `
	DELETE FROM books WHERE id=$1;
	`
	rows, err := config.DB.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "data is not found",
		})
		return
	}

	count, _ := rows.RowsAffected()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"deleted data amount": count,
	})

}
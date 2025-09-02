package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"

	// "bytes"

	"github.com/gin-gonic/gin"
)

func Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// var b bytes.Buffer
		// path := "images/covers/"
		// typer := ".png"
		// b.WriteString(path)
		// b.WriteString(book.Name)
		// b.WriteString(typer)
		query := "INSERT INTO books (name,author,price,genre,library_id,is_avaible) VALUES($1,$2,$3,$4,$5,$6)"
		_, err := database.DB.Exec(query, book.Name, book.Author, book.Price, book.Genre, book.Library_id, book.Is_avaible)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error(), "success": false})
			return
		}
		c.JSON(200, book)
	}
}

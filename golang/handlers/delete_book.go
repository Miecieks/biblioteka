package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"

	"github.com/gin-gonic/gin"
)

func Delete_book() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		query := "DELETE FROM books WHERE id = $1;"
		_, err := database.DB.Exec(query, book.Id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error(), "success": false})
			return
		}
		c.JSON(200, book)

	}
}

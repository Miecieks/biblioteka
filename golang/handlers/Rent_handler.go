package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"

	"github.com/gin-gonic/gin"
)

func Rent_book() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Rent
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// TodayDate := dateOfToday+1month

		query := "INSERT INTO rented (user_id, book_id,to_return,penalty,is_extended)"
		_, err := database.DB.Exec(query, book.User_id, book.Book_id)
		if err != nil {
			c.JSON(500, gin.H{"error": "problem z wypozyczeniem"})
		}

	}
}

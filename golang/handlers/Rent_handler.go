package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"

	// "bytes"
	"time"

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

		quer := "SELECT library_id FROM Books WHERE id = $1"
		res := database.DB.QueryRow(quer, book.Book_id)
		errn := res.Scan(&book.Library_id)
		if errn != nil {
			c.JSON(500, gin.H{"error": "problem z wypozyczeniem", "errD": errn.Error()})
		}

		quer = "SELECT penalty_per_day FROM Libraries WHERE id = $1"
		res = database.DB.QueryRow(quer, book.Library_id)
		errn = res.Scan(&book.Penalty)
		if errn != nil {
			c.JSON(500, gin.H{"error": "problem z wypozyczeniem", "errD": errn.Error()})
		}

		a := time.Now().String()
		sliced := a[0:10]
		// year := sliced[0:4]
		// month := sliced[5:7]
		// day := sliced[8:10]
		// var buffer bytes.Buffer

		// buffer.WriteString(day)
		// buffer.WriteString("-")
		// buffer.WriteString(month)
		// buffer.WriteString("-")
		// buffer.WriteString(year)

		// datte := buffer.String()
		query := "INSERT INTO rented (user_id, book_id,to_return,penalty,is_extended) VALUES($1,$2,$3,$4,$5)"
		_, err := database.DB.Exec(query, book.User_id, book.Book_id, sliced, book.Penalty, false)
		if err != nil {
			c.JSON(500, gin.H{"error": "problem z wypozyczeniem", "errD": err.Error()})
		}

	}
}

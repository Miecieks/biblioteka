package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
	"bytes"
	"strconv"

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

		TkQuery := "SELECT is_avaible FROM books WHERE id = $1"
		respo := database.DB.QueryRow(TkQuery, book.Book_id)
		var ledio bool
		eern := respo.Scan(&ledio)
		if eern != nil {
			c.JSON(500, gin.H{"error": "problem z wypozyczeniem", "errD": eern.Error()})
		}

		if ledio == false {
			c.JSON(321, gin.H{"res": "Ksiazka jest juz wypozyczona", "errD": "none"})
		} else {

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
			year := sliced[0:4]
			month := sliced[5:7]
			day := sliced[8:10]
			var buffer bytes.Buffer
			month_temp, _ := strconv.Atoi(month)
			month_temp = month_temp + 1
			month = strconv.Itoa(month_temp)

			buffer.WriteString(year)
			buffer.WriteString("-")
			buffer.WriteString(month)
			buffer.WriteString("-")
			buffer.WriteString(day)

			datte := buffer.String()
			query := "INSERT INTO rented (user_id, book_id,to_return,penalty,is_extended) VALUES($1,$2,$3,$4,$5)"
			_, err := database.DB.Exec(query, book.User_id, book.Book_id, datte, book.Penalty, false)
			if err != nil {
				c.JSON(500, gin.H{"error": "problem z wypozyczeniem", "errD": err.Error()})
			}

			query = "UPDATE books SET is_avaible = $1 WHERE id = $2"
			_, err = database.DB.Exec(query, false, book.Book_id)
			if err != nil {
				c.JSON(500, gin.H{"error": "problem z wypozyczeniem", "errD": err.Error()})
			}
			c.JSON(200, gin.H{"res": "Udalo sie!", "errD": "none"})
		}
	}
}

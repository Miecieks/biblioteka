package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.RegisterStruct
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		idik := 0
		result := database.DB.QueryRow("SELECT id FROM users ORDER BY id DESC")
		errn := result.Scan(&idik)
		if errn != nil {
			c.JSON(400, gin.H{"error": "bląd w odczycie id"})
			return
		}
		idik = idik + 1
		temp_idik := idik
		UID := ""
		var buffer bytes.Buffer
		for i := 1; i < 8; i++ {
			idik = idik / 10
			if idik == 0 {
				temp := 8 - i
				for x := 0; x < temp; x++ {
					buffer.WriteString("0")
				}
				buffer.WriteString(strconv.Itoa(temp_idik))

				UID = buffer.String()
				break
			}
		}
		fmt.Println(strconv.Itoa(temp_idik))
		fmt.Println(UID)
		query := "INSERT INTO users (user_name, password, email,first_name,last_name,user_inner_id) VALUES ($1, $2, $3,$4,$5,$6)"
		_, err := database.DB.Exec(query, data.Login, data.Pass, data.Email, data.FirstName, data.LastName, UID)
		if err != nil {
			c.JSON(500, gin.H{"error": "problem z utworzeniem usera"})
			return
		}

		c.JSON(http.StatusCreated, data)
		log.Printf("User! %s udalo sie!", data.Login)
	}
}

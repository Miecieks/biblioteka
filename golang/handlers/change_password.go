package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ChangePasses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		currentPass := ""
		res := database.DB.QueryRow("SELECT password FROM users WHERE id = $1", user.Id)
		errn := res.Scan(&currentPass)
		if errn != nil {
			c.JSON(400, gin.H{"error": "bląd w odczycie hasła"})
			return
		}

		if currentPass == user.Password {
			c.JSON(400, gin.H{"error": "hasło musi być inne od aktualnego"})
			return
		}

		querry := "UPDATE users SET password = $1 WHERE id = $2"
		_, errn2 := database.DB.Exec(querry, user.Password, user.Id)
		if errn2 != nil {
			c.JSON(400, gin.H{"error": "bląd w wykonywaniu akcji"})
			return
		}

		c.JSON(200, gin.H{"success": "Udało się!"})
		fmt.Print("powiodło sie")
	}
}

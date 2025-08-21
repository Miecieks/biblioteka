package main

import (
	"log"
	"strconv"
	"time"

	db "biblioteka-backend/database"
	"biblioteka-backend/handlers"
	"biblioteka-backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8085"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db.ConnectDatabase()
	router.GET("/ping", func(c *gin.Context) {
		ping_test, err := handlers.GetUserByID(1)
		if err != nil {
			log.Printf("Błąd podczas wykonywania zapytania: %v", err)
		}
		c.JSON(200, gin.H{
			"message": ping_test,
		})
	})
	router.GET("api/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = action[1:]
		if action == "" {
			c.JSON(400, gin.H{
				"message": "Brak parametru",
			})
		} else {
			switch action {
			case "get":
				id, err := strconv.Atoi(name)
				if err != nil {
					c.JSON(400, gin.H{"error": "Niepoprawne ID"})
					return
				}
				user, err := handlers.GetUserByID(id)
				if err != nil {
					c.JSON(404, gin.H{"error": "Użytkownik nie znaleziony"})
				} else {
					c.JSON(200, gin.H{
						"id":            user.Id,
						"first_name":    user.First_name,
						"last_name":     user.Last_name,
						"user_name":     user.User_name,
						"email":         user.Email,
						"admin":         user.Admin,
						"user_inner_id": user.User_inner_id,
						"pfp":           user.Pfp,
					})
				}
			}
		}
	})

	router.POST("api/user/:action", func(c *gin.Context) {

		action := c.Param("action")
		if action == "" {
			c.JSON(400, gin.H{
				"message": "Brak parametru",
			})
		} else {

			switch action {
			case "verify":
				var r models.LoginStruct
				if err := c.BindJSON(&r); err != nil {
					c.JSON(400, gin.H{"success": false, "res": "Niepoprawny JSON"})
					return
				}
				user, err := handlers.Verify(r.Login, r.Pass)
				if err != nil {
					c.JSON(404, gin.H{"success": false, "res": "Błąd podczas weryfikacji"})
				} else {
					c.JSON(200, gin.H{
						"success": true,
						"res":     user,
					})
				}

			}

		}
	})

	router.POST("api/user/add", handlers.Register())

	router.Run()
}

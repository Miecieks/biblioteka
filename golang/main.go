package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	db "biblioteka-backend/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "qqqq")
}

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

	router.GET("/ping", func(c *gin.Context) {
		ping_test, err := db.Initialize("SELECT user_name FROM users")
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
				query := "SELECT * FROM users WHERE id = '" + name + "'"
				qqq, err := db.Initialize(query)
				if err != nil {
					log.Printf("Błąd podczas wykonywania zapytania: %v", err)
				}
				c.JSON(200, gin.H{
					"veryf":     "true",
					"FirstName": qqq[1],
					"LastName":  qqq[2],
					"username":  qqq[3],
					"email":     qqq[5],
					"admin":     qqq[6],
					"UID":       qqq[7],
					"pfp":       qqq[8],
				})
			}
		}
	})

	type LoginStruct struct {
		Login string `json:"login"`
		Pass  string `json:"pass"`
	}

	router.POST("api/user/:action", func(c *gin.Context) {
		var r LoginStruct
		action := c.Param("action")
		if action == "" {
			c.JSON(400, gin.H{
				"message": "Brak parametru",
			})
		} else {

			switch action {
			case "verify":
				if err := c.BindJSON(&r); err != nil {
					c.JSON(400, gin.H{"success": false, "res": "Niepoprawny JSON"})
					return
				}
				query := "SELECT * FROM users WHERE user_name = '" + r.Login + "' AND password = '" + r.Pass + "'"
				qqq, err := db.Initialize(query)
				if err != nil {
					log.Printf("Błąd podczas wykonywania zapytania SQL: %v", err)
					c.JSON(401, gin.H{"failed": true, "res": "Błąd logowania"})
				} else if qqq != nil {
					tP := fmt.Sprintf("%d", qqq[0])
					log.Printf("Jest gicior majonez dla:"+tP, err)
					c.JSON(200, gin.H{"success": true, "message": "Zalogowano!", "id": qqq[0]})
				} else {
					log.Printf("nie jest gicior", err)
					c.JSON(200, gin.H{"success": false, "message": "nie zalogowano"})
				}
			}

		}
	})

	router.Run()
}

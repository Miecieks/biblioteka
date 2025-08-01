package main

import (
	passes "biblioteka-backend/packages"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "qqqq")
}

func initialize(query string) ([]interface{}, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, passes.Certs("nazwa"))
	if err != nil {
		return nil, fmt.Errorf("błąd połączenia z bazą: %v", err)
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("błąd zapytania: %v", err)
	}
	defer rows.Close()

	var results []interface{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("błąd pobierania wartości: %v", err)
		}
		results = append(results, values...)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("błąd podczas iteracji po wierszach: %v", rows.Err())
	}

	return results, nil
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
		ping_test, err := initialize("SELECT user_name FROM users")
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
				query := "SELECT * FROM users WHERE user_name = '" + name + "'"
				qqq, err := initialize(query)
				if err != nil {
					log.Printf("Błąd podczas wykonywania zapytania: %v", err)
				}
				c.JSON(200, gin.H{
					"veryf":     "true",
					"FirstName": qqq[1],
					"LastName":  qqq[2],
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
				qqq, err := initialize(query)
				if err != nil {
					log.Printf("Błąd podczas wykonywania zapytania SQL: %v", err)
				}
				c.JSON(200, gin.H{"success": true, "message": "Zalogowano!", "id": qqq[0]})

			}

		}
	})

	router.Run()
}

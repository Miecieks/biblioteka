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
						"password":      user.Password,
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
	router.POST("api/user/passChange", handlers.ChangePasses())
	router.POST("api/stuff/libraries/book/add", handlers.Insert())
	router.POST("api/stuff/libraries/book/delete", handlers.Delete_book())
	router.POST("api/stuff/libraries/book/rent", handlers.Rent_book())

	router.GET("api/stuff/libraries/get/all", func(c *gin.Context) {
		libr := handlers.GetAllLibraries()
		c.JSON(200, gin.H{
			"success": true,
			"res":     libr,
		})

	})
	router.GET("api/stuff/libraries/book/get", func(c *gin.Context) {
		res := handlers.GetAllBooks()
		c.JSON(200, gin.H{
			"success": true,
			"res":     res,
		})

	})

	router.GET("api/stuff/libraries/book/get/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, gin.H{
				"message": "Brak parametru",
				"success": false,
			})
		} else {
			id_int, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "błąd w konwersji",
					"success": false,
				})
			}
			res, err := handlers.GetBookById(id_int)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "błąd w konwersji",
					"success": false,
				})
			}
			c.JSON(200, gin.H{
				"success": true,
				"res":     res,
			})
		}
	})
	router.GET("api/stuff/libraries/book/get/:id/extra", func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, gin.H{
				"message": "Brak parametru",
				"success": false,
			})
		} else {
			id_int, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "błąd w konwersji",
					"success": false,
				})
			}
			res, err := handlers.GetBookByIdWithLibrary(id_int)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "błąd w konwersji",
					"success": false,
				})
			}
			c.JSON(200, gin.H{
				"success": true,
				"res":     res,
			})
		}
	})

	router.Run()
}

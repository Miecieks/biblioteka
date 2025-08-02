package handlers

import (
	passes "biblioteka-backend/packages"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		conn, err := pgx.Connect(ctx, passes.Certs("nazwa"))
		if err != nil {
			fmt.Errorf("błąd połączenia z bazą: %v", err)
		}
		defer conn.Close(ctx)

		c.JSON(http.StatusOK, gin.H{"succeded": true, "message": "zarejestrowano"})
	}
}

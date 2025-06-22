package main

import (
	passes "biblioteka-backend/packages"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, passes.Certs("nazwa"))
	if err != nil {
		log.Fatalf("Błąd połączenia z bazą: %v", err)
	}
	defer conn.Close(ctx)
	var currentTime string
	err = conn.QueryRow(ctx, "SELECT user_name FROM users").Scan(&currentTime)
	if err != nil {
		log.Fatalf("Błąd zapytania: %v", err)
	}
	fmt.Println("Aktualny czas z bazy:", currentTime)
}

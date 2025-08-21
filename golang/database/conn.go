package database

import (
	"database/sql"
	"fmt"
	"log"

	passes "biblioteka-backend/packages"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	dsn := passes.Certs("nazwa")
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	fmt.Println("Database connected!")
}

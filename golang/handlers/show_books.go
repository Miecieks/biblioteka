package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
	"fmt"
)

func GetAllBooks() []models.Book {

	result, err := database.DB.Query("SELECT * FROM books")
	if err != nil {
		fmt.Print("błąd z żądaniem", err.Error())
		return nil
	}

	books := []models.Book{}
	for result.Next() {
		var book models.Book
		err = result.Scan(&book.Id, &book.Name, &book.Author, &book.Price, &book.Genre, &book.Library_id, &book.Is_avaible, &book.Cover)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}

	return books
}

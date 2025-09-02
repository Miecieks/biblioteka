package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
)

func GetBookById(id int) (*models.Book, error) {
	var book models.Book
	row := database.DB.QueryRow("SELECT * FROM books WHERE id = $1", id)
	err := row.Scan(&book.Id, &book.Name, &book.Author, &book.Price, &book.Genre, &book.Library_id, &book.Is_avaible, &book.Cover)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

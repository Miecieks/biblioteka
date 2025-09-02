package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
)

func GetBookByIdWithLibrary(id int) (*models.BookWithLib, error) {
	var book models.BookWithLib
	row := database.DB.QueryRow("SELECT books.id,books.name,books.author,books.price,books.genre,books.library_id,books.is_avaible,books.cover,Libraries.name,Libraries.location FROM books INNER JOIN Libraries ON books.library_id = Libraries.id WHERE books.id = $1", id)
	err := row.Scan(&book.Id, &book.Name, &book.Author, &book.Price, &book.Genre, &book.Library_id, &book.Is_avaible, &book.Cover, &book.LibName, &book.Location)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

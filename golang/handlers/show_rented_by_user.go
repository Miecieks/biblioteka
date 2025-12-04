package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
)

func ShowRentedOne(id int) (*models.Rent, error) {
	var rent models.Rent
	row := database.DB.QueryRow("SELECT * FROM rented WHERE user_id = $1;", id)
	err := row.Scan(&rent.Id, &rent.User_id, &rent.Book_id, &rent.To_return, &rent.Penalty, &rent.Is_extended)
	if err != nil {
		return nil, err
	}
	return &rent, nil
}

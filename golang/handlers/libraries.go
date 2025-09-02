package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
	"fmt"
)

func GetAllLibraries() []models.Library {

	result, err := database.DB.Query("SELECT * FROM Libraries")
	if err != nil {
		fmt.Print("błąd z żądaniem", err.Error())
		return nil
	}

	librarium := []models.Library{}
	for result.Next() {
		var library models.Library
		err = result.Scan(&library.Id, &library.Name, &library.Location, &library.Penalty_per_day)
		if err != nil {
			panic(err.Error())
		}
		librarium = append(librarium, library)
	}
	return librarium
}

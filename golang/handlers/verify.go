package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
)

func Verify(login string, password string) (*models.ResForVerify, error) {
	var logger models.ResForVerify
	row := database.DB.QueryRow("SELECT id, admin FROM users WHERE user_name = $1 AND password = $2", login, password)
	err := row.Scan(&logger.ID, &logger.Admin)
	if err != nil {
		return nil, err
	}
	return &logger, nil
}

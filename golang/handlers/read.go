package handlers

import (
	"biblioteka-backend/database"
	"biblioteka-backend/models"
)

func GetUserByID(id int) (*models.User, error) {
	var user models.User
	row := database.DB.QueryRow("SELECT id, first_name, last_name,user_name, password ,email,admin,user_inner_id,pfp FROM users WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.First_name, &user.Last_name, &user.User_name, &user.Password, &user.Email, &user.Admin, &user.User_inner_id, &user.Pfp)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

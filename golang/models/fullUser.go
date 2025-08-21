package models

type User struct {
	Id            int
	First_name    string
	Last_name     string
	User_name     string
	Password      string
	Email         string
	Admin         bool
	User_inner_id string
	Pfp           string
}

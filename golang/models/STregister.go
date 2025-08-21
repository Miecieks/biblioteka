package models

type RegisterStruct struct {
	Login     string `json:"login"`
	Pass      string `json:"pass"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

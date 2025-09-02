package models

type Book struct {
	Id         int
	Name       string
	Author     string
	Price      float32
	Genre      string
	Library_id int
	Is_avaible bool
	Cover      string
}

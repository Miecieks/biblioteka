package models

type Rent struct {
	Id      int
	User_id int
	Book_id int
	// To_return   date
	Penalty     float32
	Is_extended bool
}

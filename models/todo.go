package models

//Todo -
type Todo struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Texto  string `json:"texto"`
	DtTodo string `json:"dtTodo"`
	IDUser int    `json:"idUser"`
}

package models

//Todo - estrutura para inserir um todo
type Todo struct {
	ID      int     `json:"id"`
	Titulo  string  `json:"titulo"`
	Texto   string  `json:"texto"`
	DtTodo  string  `json:"dtTodo"`
	Usuario Usuario `json:"usuario"`
}

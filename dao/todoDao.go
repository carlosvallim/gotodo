package dao

import (
	"fmt"

	"github.com/carlosvallim/gotodo/data"
	"github.com/carlosvallim/gotodo/models"
)

//FindTodos - busca todas as anotações
func FindTodos() ([]*models.Todo, error) {

	defer func() {
		if ve := recover(); ve != nil {
			err := fmt.Errorf("%v", ve)
			fmt.Printf("Erro: %v\n", err)
		}
	}()

	rows, err := data.Db.Query(`
		select td.id, 
			   td.titulo, 
			   td.texto, 
			   td.dt_todo,
			   td.id_user, 
			   us.nome
		from todo td
		  inner join usuario us on (us.id = td.id_user)
		order by us.nome		  
	`)
	if err != nil {
		panic(fmt.Errorf("erro ao listar os todos: %v", err))
	}

	todos := []*models.Todo{}

	for rows.Next() {
		todo := &models.Todo{}
		err := rows.Scan(&todo.ID, &todo.Titulo, &todo.Texto, &todo.DtTodo, &todo.Usuario.ID, &todo.Usuario.Nome)
		if err != nil {
			panic(fmt.Errorf("erro ao efetuar scan dos todos: %v", err))
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

package gotodo

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/carlosvallim/gotodo/dao"
	"github.com/carlosvallim/gotodo/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	todos []*models.Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*models.Todo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) FindTodos(ctx context.Context) ([]*models.Todo, error) {
	todos, err := dao.FindTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

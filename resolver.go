package gotodo

import (
	"context"

	"github.com/carlosvallim/gotodo/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	todos []*Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*models.Todo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) FindTodos(ctx context.Context) ([]*models.Todo, error) {
	panic("not implemented")
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) Usuario(ctx context.Context, obj *models.Todo) (*Usuario, error) {
	panic("not implemented")
}

//go:generate go run scripts/gqlgen.go -v

package gqlgen_sample

import (
	"context"
	"fmt"
	"go_graphql/gqlgen_sample/gettingstarted"
	"math/rand"
)

type Resolver struct {
	todos []gqlgen_sample.Todo
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

func (r *mutationResolver) CreateToDo(ctx context.Context, input NewTodo) (gqlgen_sample.Todo, error) {
	todo := gqlgen_sample.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]gqlgen_sample.Todo, error) {
	return r.todos, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *gqlgen_sample.Todo) (User, error) {
	return User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

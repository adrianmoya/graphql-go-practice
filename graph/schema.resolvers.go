package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/adrianmoya/graphql-go-practice/graph/generated"
	"github.com/adrianmoya/graphql-go-practice/graph/model"
	"github.com/adrianmoya/graphql-go-practice/jwt"
	"github.com/adrianmoya/graphql-go-practice/middleware"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID, // fix this line
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) Login(ctx context.Context, username *string, password *string) (*model.LoginOutput, error) {
	if *username == "gocommunity" && *password == "endava2021" {
		token, err := jwt.CreateToken(*username)
		return &model.LoginOutput{
			Token: token,
		}, err
	}
	return nil, errors.New("Unauthorized")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	user := middleware.ForContext(ctx)
	if user == "" {
		return nil, errors.New("Unauthorized")
	}
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

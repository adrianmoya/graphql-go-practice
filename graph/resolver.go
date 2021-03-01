package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/adrianmoya/graphql-go-practice/graph/model"

type Resolver struct {
	todos []*model.Todo
}

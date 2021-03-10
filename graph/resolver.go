package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/3dw1nM0535/galva/graph/model"
)

type Resolver struct {
	todos []*model.Todo
}

func New() *Resolver {
	todos := make([]*model.Todo, 0)
	todos = append(todos, &model.Todo{ID: "1", Text: "deploy some code", Done: false})
	todos = append(todos, &model.Todo{ID: "2", Text: "write some tests", Done: false})

	return &Resolver{
		todos: todos,
	}
}

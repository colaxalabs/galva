package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/3dw1nM0535/galva/store"
)

// Resolver returns the entry point
type Resolver struct {
	ORM *store.ORM
}

// New returns a new resolver
func New(orm *store.ORM) *Resolver {
	return &Resolver{
		ORM: orm,
	}
}

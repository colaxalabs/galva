package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/store/models"
)

func (r *landResolver) ID(ctx context.Context, obj *models.Land) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *landResolver) Size(ctx context.Context, obj *models.Land) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Land returns generated.LandResolver implementation.
func (r *Resolver) Land() generated.LandResolver { return &landResolver{r} }

type landResolver struct{ *Resolver }

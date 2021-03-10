package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/store/models"
)

func (r *landResolver) ID(ctx context.Context, obj *models.Land) (string, error) {
	id := obj.ID.String()
	return id, nil
}

// Land returns generated.LandResolver implementation.
func (r *Resolver) Land() generated.LandResolver { return &landResolver{r} }

type landResolver struct{ *Resolver }

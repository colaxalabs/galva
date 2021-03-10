package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/graph/model"
	"github.com/3dw1nM0535/galva/store/models"
)

func (r *mutationResolver) CreateLand(ctx context.Context, input model.NewLand) (*models.Land, error) {
	newLand := &models.Land{
		ID:            input.TokenID,
		PostalCode:    input.PostalCode,
		SateliteImage: input.SateliteImage,
		State:         input.State,
		Location:      input.Location,
	}
	r.ORM.Store.Create(&newLand)
	return newLand, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

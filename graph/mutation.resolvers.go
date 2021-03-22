package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/graph/model"
	"github.com/3dw1nM0535/galva/store/models"
	"github.com/3dw1nM0535/galva/utils"
)

func (r *mutationResolver) AddUser(ctx context.Context, input model.RegisterUser) (*models.User, error) {
	// Check if user already exists
	parsedAddress := utils.ParseAddress(input.Address)
	user := &models.User{}
	r.ORM.Store.Where("address = ?", parsedAddress).First(&user)
	if user.ID != nil {
		return nil, errors.New("user already exists")
	}
	id := models.NewID()
	newUser := &models.User{
		ID:        id,
		Address:   parsedAddress,
		Signature: input.Signature,
	}
	r.ORM.Store.Save(&newUser)
	return newUser, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

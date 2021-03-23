package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/3dw1nM0535/galva/constants"
	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/store/models"
)

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello", nil
}

func (r *queryResolver) GetProperty(ctx context.Context, id string) (*models.Property, error) {
	property := &models.Property{}
	// Validate if property exists
	r.ORM.Store.Where("id = ?", id).First(&property)
	if property.ID == 0 {
		return nil, errors.New(constants.NonExistentProperty)
	}
	return property, nil
}

func (r *queryResolver) GetUser(ctx context.Context, address string) (*models.User, error) {
	user := &models.User{}
	// Validate if user exists with us
	r.ORM.Store.Where("address = ?", address).First(&user)
	if user.ID == nil {
		return nil, errors.New(constants.NonExistentUser)
	}
	r.ORM.Store.Where("address = ?", address).First(&user)
	return user, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

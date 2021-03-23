package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/store/models"
	"github.com/3dw1nM0535/galva/utils"
)

func (r *offerResolver) ID(ctx context.Context, obj *models.Offer) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *offerResolver) User(ctx context.Context, obj *models.Offer) (*models.User, error) {
	user := &models.User{}
	userAddress := utils.ParseAddress(obj.UserAddress)
	r.ORM.Store.Where("address = ?", userAddress).Find(&user)
	return user, nil
}

func (r *offerResolver) Property(ctx context.Context, obj *models.Offer) (*models.Property, error) {
	property := &models.Property{}
	id := obj.PropertyID
	r.ORM.Store.Where("id = ?", id).Find(&property)
	return property, nil
}

func (r *propertyResolver) User(ctx context.Context, obj *models.Property) (*models.User, error) {
	user := &models.User{}
	userAddress := utils.ParseAddress(obj.UserAddress)
	r.ORM.Store.Where("address = ?", userAddress).Find(&user)
	return user, nil
}

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	id := obj.ID.String()
	return id, nil
}

// Offer returns generated.OfferResolver implementation.
func (r *Resolver) Offer() generated.OfferResolver { return &offerResolver{r} }

// Property returns generated.PropertyResolver implementation.
func (r *Resolver) Property() generated.PropertyResolver { return &propertyResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type offerResolver struct{ *Resolver }
type propertyResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *offerResolver) Title(ctx context.Context, obj *models.Offer) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

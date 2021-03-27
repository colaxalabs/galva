package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/store/models"
	"github.com/3dw1nM0535/galva/utils"
)

func (r *offerResolver) ID(ctx context.Context, obj *models.Offer) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *offerResolver) Size(ctx context.Context, obj *models.Offer) (string, error) {
	size := obj.Size.String()
	return size, nil
}

func (r *offerResolver) Cost(ctx context.Context, obj *models.Offer) (string, error) {
	cost := obj.Cost.String()
	return cost, nil
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

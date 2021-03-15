package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/store/models"
	"github.com/3dw1nM0535/galva/utils"
)

func (r *landResolver) LandOwner(ctx context.Context, obj *models.Land) (*models.User, error) {
	user := &models.User{}
	parsedAddress := utils.ParseAddress(obj.UserAddress)
	r.ORM.Store.Where("address = ?", parsedAddress).Find(&user)
	return user, nil
}

func (r *landResolver) LandOffers(ctx context.Context, obj *models.Land) ([]*models.Offer, error) {
	offers := []*models.Offer{}
	id := obj.ID
	r.ORM.Store.Where("land_id = ?", id).Find(&offers)
	return offers, nil
}

func (r *offerResolver) ID(ctx context.Context, obj *models.Offer) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *offerResolver) OfferOwner(ctx context.Context, obj *models.Offer) (*models.User, error) {
	user := &models.User{}
	address := utils.ParseAddress(obj.Tenant)
	r.ORM.Store.Where("address = ?", address).Find(&user)
	return user, nil
}

func (r *offerResolver) Land(ctx context.Context, obj *models.Offer) (*models.Land, error) {
	land := &models.Land{}
	id := obj.LandID
	r.ORM.Store.Where("id = ?", id).Find(&land)
	return land, nil
}

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *userResolver) Properties(ctx context.Context, obj *models.User) ([]*models.Land, error) {
	lands := []*models.Land{}
	parsedAddress := utils.ParseAddress(obj.Address)
	r.ORM.Store.Where("user_address = ?", parsedAddress).Find(&lands)
	return lands, nil
}

func (r *userResolver) UserOffers(ctx context.Context, obj *models.User) ([]*models.Offer, error) {
	offers := []*models.Offer{}
	address := utils.ParseAddress(obj.Address)
	r.ORM.Store.Where("tenant = ?", address).Find(&offers)
	return offers, nil
}

// Land returns generated.LandResolver implementation.
func (r *Resolver) Land() generated.LandResolver { return &landResolver{r} }

// Offer returns generated.OfferResolver implementation.
func (r *Resolver) Offer() generated.OfferResolver { return &offerResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type landResolver struct{ *Resolver }
type offerResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/graph/model"
	"github.com/3dw1nM0535/galva/store/models"
	"github.com/3dw1nM0535/galva/utils"
)

func (r *mutationResolver) CreateLand(ctx context.Context, input model.NewLand) (*models.Land, error) {
	// Check that there is such existing user
	targetUser := &models.User{}
	parsedAddress := utils.ParseAddress(input.Address)
	r.ORM.Store.Where("address = ?", parsedAddress).Find(&targetUser)
	if targetUser.ID == nil {
		return nil, fmt.Errorf("user account %s cannot be found", parsedAddress)
	}
	// Check that there is no such existing property
	targetLand := &models.Land{}
	r.ORM.Store.Where("id = ?", input.TokenID).Find(&targetLand)
	if targetLand.ID != 0 {
		return nil, fmt.Errorf("duplicate land record %d", input.TokenID)
	}
	// Proceed to creating property
	newLand := &models.Land{
		ID:            input.TokenID,
		PostalCode:    input.PostalCode,
		SateliteImage: input.SateliteImage,
		State:         input.State,
		Location:      input.Location,
		UserAddress:   utils.ParseAddress(input.Address),
		LandOwner:     targetUser,
	}
	r.ORM.Store.Save(&newLand)
	return newLand, nil
}

func (r *mutationResolver) AddUser(ctx context.Context, input model.RegisterUser) (*models.User, error) {
	user := &models.User{}
	parsedAddress := utils.ParseAddress(input.Address)
	r.ORM.Store.Where("address = ?", parsedAddress).Find(&user)
	if user.ID != nil {
		return nil, fmt.Errorf("duplicate user account: %s", parsedAddress)
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

func (r *mutationResolver) CreateOffer(ctx context.Context, input model.OfferInput) (*models.Offer, error) {
	parsedTenant := utils.ParseAddress(input.Tenant)
	parsedOwner := utils.ParseAddress(input.Owner)
	// Check if user is authenticated
	user := &models.User{}
	r.ORM.Store.Where("address = ?", parsedOwner).Find(&user)
	if user.ID == nil {
		return nil, fmt.Errorf("user %v is not authenticated", parsedOwner)
	}
	tenant := &models.User{}
	r.ORM.Store.Where("address = ?", parsedTenant).Find(&tenant)
	if tenant.ID == nil {
		return nil, fmt.Errorf("user %v is not authenticated", parsedTenant)
	}
	// Ensure property owner does not make offer to him/herself
	// Check if property exists
	land := &models.Land{}
	r.ORM.Store.Where("id = ?", input.TokenID).Find(&land)
	if land.ID == 0 {
		return nil, fmt.Errorf("unable to find property with id %v", input.TokenID)
	}
	if land.UserAddress == parsedTenant {
		return nil, fmt.Errorf("forbidden to make offer to property %v", input.TokenID)
	}
	if land.State != "Leasing" {
		return nil, fmt.Errorf("property %v is not leasing", input.TokenID)
	}
	// Proceed to creating offer
	id := models.NewID()
	formattedTime, _ := utils.ParseTime(int64(input.Duration))
	newOffer := &models.Offer{
		ID:         id,
		LandID:     land.ID,
		Purpose:    input.Purpose,
		Size:       input.Size,
		Duration:   formattedTime,
		Cost:       input.Cost,
		Owner:      parsedOwner,
		Tenant:     parsedTenant,
		Title:      input.Title,
		FullFilled: false,
	}
	r.ORM.Store.Save(&newOffer)
	return newOffer, nil
}

func (r *mutationResolver) ChangeState(ctx context.Context, input model.StateInput) (*models.Land, error) {
	land := &models.Land{}
	r.ORM.Store.Where("id = ?", input.ID).Find(&land)
	if land == nil {
		return nil, fmt.Errorf("unable to find property %v", input.ID)
	}
	land.State = input.State
	r.ORM.Store.Save(&land)
	return land, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

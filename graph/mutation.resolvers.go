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
		User:          targetUser,
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

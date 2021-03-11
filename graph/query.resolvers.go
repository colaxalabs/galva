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

func (r *queryResolver) Lands(ctx context.Context) ([]*models.Land, error) {
	lands := []*models.Land{}
	r.ORM.Store.Preload("User").Find(&lands)
	return lands, nil
}

func (r *queryResolver) NearByPostal(ctx context.Context, postal string) ([]*models.Land, error) {
	lands := []*models.Land{}
	r.ORM.Store.Where("postal_code = ? AND state = ?", postal, "Leasing").Preload("User").Find(&lands)
	return lands, nil
}

func (r *queryResolver) User(ctx context.Context, address string) (*models.User, error) {
	user := &models.User{}
	parsedAddress := utils.ParseAddress(address)
	r.ORM.Store.Where("address = ?", parsedAddress).Preload("Properties").Find(&user)
	if user.ID == nil {
		return nil, fmt.Errorf("account %s cannot be found", parsedAddress)
	}
	return user, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

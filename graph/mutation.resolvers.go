package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/3dw1nM0535/galva/eth"
	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/graph/model"
	"github.com/3dw1nM0535/galva/nft"
	"github.com/3dw1nM0535/galva/store/models"
	"github.com/3dw1nM0535/galva/utils"
	"github.com/ethereum/go-ethereum/common"
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

func (r *mutationResolver) AddListing(ctx context.Context, input model.PropertyInput) (*models.Property, error) {
	property := &models.Property{}
	// Validate the property is tokenized
	newEth, err := eth.NewEthClient()
	if err != nil {
		return nil, fmt.Errorf("Error '%v' while setting up ethereum node", err)
	}
	// Load contract to query nft
	contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	instance, err := nft.NewNft(contractAddress, newEth.Client)
	if err != nil {
		return nil, fmt.Errorf("Error '%v' while setting up ethereum contract", err)
	}
	parsedPropertyId, _ := strconv.ParseInt(string(input.ID), 10, 64)
	tokenId := big.NewInt(parsedPropertyId)
	owner, err := instance.OwnerOf(nil, tokenId)
	if err != nil {
		return nil, fmt.Errorf("Error '%v' querying token owner", err)
	}
	parsedOwner := owner.String()
	// Create new listing
	newListing := &models.Property{
		ID:            parsedPropertyId,
		PostalCode:    input.PostalCode,
		Location:      input.Location,
		SateliteImage: input.SateliteImage,
		UserAddress:   parsedOwner,
	}
	r.ORM.Store.Save(&newListing)
	return property, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/3dw1nM0535/galva/constants"
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
	// Load contract to query nft validity
	contractAddress := common.HexToAddress(constants.NftContractAddress)
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

func (r *mutationResolver) MakeOffer(ctx context.Context, input model.OfferInput) (*models.Offer, error) {
	property := &models.Property{}
	// Validate the property is tokenized
	newEth, err := eth.NewEthClient()
	if err != nil {
		return nil, fmt.Errorf("Error '%v' while setting up ethereum node", err)
	}
	// Load contract to query nft validity
	contractAddress := common.HexToAddress(constants.NftContractAddress)
	instance, err := nft.NewNft(contractAddress, newEth.Client)
	if err != nil {
		return nil, fmt.Errorf("Error '%v' while setting up ethereum contract", err)
	}
	// Check if property id is valid
	propertyId := big.NewInt(int64(input.PropertyID))
	owner, err := instance.OwnerOf(nil, propertyId)
	if err != nil {
		return nil, fmt.Errorf("Error '%v' querying token owner", err)
	}
	// Check if property is listed
	r.ORM.Store.Where("id = ?", input.PropertyID).First(&property)
	if property.ID == 0 {
		return nil, fmt.Errorf("cannot find listing in market with id %v", input.PropertyID)
	}
	parsedOwner := owner.String()
	// Create offer to property
	id := models.NewID()
	newOffer := &models.Offer{
		ID:          id,
		Purpose:     input.Purpose,
		Size:        input.Size,
		Duration:    input.Duration,
		Cost:        input.Cost,
		Owner:       parsedOwner,
		UserAddress: input.UserAddress,
		ExpiresIn:   time.Now().Add(time.Hour * 24 * 7),
	}
	return newOffer, nil
}

func (r *mutationResolver) AcceptOffer(ctx context.Context, input model.AcceptOfferInput) (*models.Offer, error) {
	// Check that the offer exists
	offer := &models.Offer{}
	r.ORM.Store.Where("id = ?", input.ID).First(&offer)
	if offer.ID == nil {
		return nil, fmt.Errorf("cannot find offer with id %v", input.ID)
	}
	// Validate that the offer owner is the requestor
	if offer.Owner != utils.ParseAddress(input.UserAddress) {
		return nil, errors.New(constants.ForbiddenToOwner)
	}
	// Offer should not be expired
	if offer.ExpiresIn.Unix() < time.Now().Unix() {
		return nil, errors.New(constants.Expired)
	}
	// Update offer
	offer.Accepted = true
	r.ORM.Store.Save(&offer)
	return offer, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
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
	// Check duplicate user account
	user := &models.User{}
	parsedAddress := utils.ParseAddress(input.Address)
	r.ORM.Store.Where("address = ?", parsedAddress).Find(&user)
	if user.ID != nil {
		return nil, fmt.Errorf("duplicate user account: %s", parsedAddress)
	}
	// Check non-duplicate account signature
	signedAccount := &models.User{}
	r.ORM.Store.Where("signature = ?", input.Signature).Find(&signedAccount)
	if signedAccount.Signature == input.Signature {
		return nil, fmt.Errorf("cannot authenticate user %v with duplicate signature for user id %v", parsedAddress, signedAccount.Address)
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
	if land.ID == 0 {
		return nil, fmt.Errorf("unable to find property %v", input.ID)
	}
	land.State = input.State
	r.ORM.Store.Save(&land)
	return land, nil
}

func (r *mutationResolver) TenantSigns(ctx context.Context, input model.SigningInput) (*models.Offer, error) {
	parsedSigner := utils.ParseAddress(input.Signer)
	// Find signer
	signer := &models.User{}
	r.ORM.Store.Where("address = ?", parsedSigner).Find(&signer)
	if signer.ID == nil {
		return nil, fmt.Errorf("unable to find signer %v", parsedSigner)
	}
	// Authenticate signer
	if signer.Signature != input.Signature {
		return nil, fmt.Errorf("unable to authenticate signature for signer %v", parsedSigner)
	}
	// Find offer
	offer := &models.Offer{}
	r.ORM.Store.Where("id = ?", input.ID).Find(&offer)
	if offer.ID == nil {
		return nil, fmt.Errorf("unable to find offer with id %v for signing", input.ID)
	}
	// Tenant cannot sign rejected offer
	if offer.Rejected {
		return nil, fmt.Errorf("forbidden to sign rejected offer %v", input.ID)
	}
	if !offer.Accepted {
		return nil, errors.New("offer not accepted by property owner")
	}
	// Only offer creator can sign the offer
	if offer.Tenant != parsedSigner {
		return nil, fmt.Errorf("forbidden to sign offer created by %v", offer.Tenant)
	}
	offer.TenantSignature = input.Signature
	r.ORM.Store.Save(&offer)
	return offer, nil
}

func (r *mutationResolver) OwnerSigns(ctx context.Context, input model.SigningInput) (*models.Offer, error) {
	parsedSigner := utils.ParseAddress(input.Signer)
	// Find signer
	signer := &models.User{}
	r.ORM.Store.Where("address = ?", parsedSigner).Find(&signer)
	if signer.ID == nil {
		return nil, fmt.Errorf("unable to find signer %v", parsedSigner)
	}
	// Authenticate signer
	if signer.Signature != input.Signature {
		return nil, fmt.Errorf("unable to authenticate signature for signer %v", parsedSigner)
	}
	// Find offer
	offer := &models.Offer{}
	r.ORM.Store.Where("id = ?", input.ID).Find(&offer)
	if offer.ID == nil {
		return nil, fmt.Errorf("unable to find offer with id %v for signing", input.ID)
	}
	// Owner cannot sign !accepted offer
	if !offer.Accepted {
		return nil, errors.New("accept offer first and tenant signature for you to sign last")
	}
	if offer.Rejected {
		return nil, errors.New("cannot sign rejected offer")
	}
	// Owner cannot sign before tenant signs
	if offer.Owner == parsedSigner && offer.TenantSignature == "" {
		return nil, fmt.Errorf("forbidden to sign offer %v before tenant", input.ID)
	}
	// Authenticate owner in the offer is the signer
	if offer.Owner != parsedSigner {
		return nil, fmt.Errorf("forbidden to sign offer where signer is not %v", offer.Owner)
	}
	offer.OwnerSignature = input.Signature
	offer.Signed = true
	r.ORM.Store.Save(&offer)
	return offer, nil
}

func (r *mutationResolver) RejectOffer(ctx context.Context, input model.OfferStateInput) (*models.Offer, error) {
	// Find offer
	offer := &models.Offer{}
	r.ORM.Store.Where("id = ?", input.ID).First(&offer)
	if offer.ID == nil {
		return nil, fmt.Errorf("unable to find offer %v", input.ID)
	}
	// Restrict rejecting already signed offer
	if offer.Signed || offer.TenantSignature != "" {
		return nil, errors.New("forbidden to reject signed offer")
	}
	// Restrict rejecting offer to property owner
	if offer.Owner != input.Creator {
		return nil, fmt.Errorf("forbidden to reject offer for user %v", offer.Owner)
	}
	if offer.Accepted {
		return nil, errors.New("forbidden to reject accepted offer")
	}
	offer.Rejected = true
	r.ORM.Store.Save(&offer)
	return offer, nil
}

func (r *mutationResolver) AcceptOffer(ctx context.Context, input model.OfferStateInput) (*models.Offer, error) {
	// Find offer
	offer := &models.Offer{}
	r.ORM.Store.Where("id = ?", input.ID).First(&offer)
	if offer.ID == nil {
		return nil, fmt.Errorf("unable to find offer %v", input.ID)
	}
	// Restrict accepting offer to property owner
	if offer.Owner != input.Creator {
		return nil, fmt.Errorf("forbidden to accept offer for user %v", offer.Owner)
	}
	// Restrict rejecting already signed offer
	if offer.Signed || offer.TenantSignature != "" {
		return nil, errors.New("forbidden to reject signed offer")
	}
	// Restrict accepting !rejected offers
	if offer.Rejected {
		return nil, errors.New("forbidden to accept rejected offer")
	}
	offer.Accepted = true
	r.ORM.Store.Save(&offer)
	return offer, nil
}

func (r *mutationResolver) DraftOffer(ctx context.Context, id string) (*models.Offer, error) {
	// Find offer
	offer := &models.Offer{}
	r.ORM.Store.Where("id = ?", id).First(&offer)
	if offer.ID == nil {
		return nil, fmt.Errorf("unable to find offer %v", id)
	}
	// Restrict drafting not signed offer
	if !offer.Signed {
		return nil, errors.New("forbidden to draft unsigned offer")
	}
	// Restrict drafting an already drafted offer
	if offer.Drafted {
		return nil, errors.New("cannot draft an already drafted offer")
	}
	offer.Drafted = true
	r.ORM.Store.Save(&offer)
	return offer, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

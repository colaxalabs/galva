// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewLand struct {
	TokenID       int    `json:"tokenId"`
	PostalCode    string `json:"postalCode"`
	SateliteImage string `json:"sateliteImage"`
	State         string `json:"state"`
	Address       string `json:"address"`
	Location      string `json:"location"`
}

type OfferInput struct {
	TokenID    int    `json:"tokenId"`
	Purpose    string `json:"purpose"`
	Size       string `json:"size"`
	Duration   int    `json:"duration"`
	Cost       string `json:"cost"`
	Owner      string `json:"owner"`
	Tenant     string `json:"tenant"`
	Title      string `json:"title"`
	FullFilled bool   `json:"fullFilled"`
}

type RegisterUser struct {
	Address   string `json:"address"`
	Signature string `json:"signature"`
}

type StateInput struct {
	ID    int    `json:"id"`
	State string `json:"state"`
}

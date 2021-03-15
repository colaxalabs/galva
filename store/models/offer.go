package models

import (
	"time"
)

type Offer struct {
	ID              *ID       `gorm:"primary_key;not_null"`
	LandID          int       `gorm:"not_null"`
	Purpose         string    `gorm:"text;not_null"`
	Size            string    `gorm:"not_null"`
	Duration        time.Time `gorm:"not_null"`
	Cost            string    `gorm:"not_null"`
	Owner           string    `gorm:"not_null"`
	Tenant          string    `gorm:"not_null"`
	Title           string    `gorm:"not_null"`
	FullFilled      bool      `gorm:"default:false;not_null"`
	OfferOwner      *User
	Land            *Land
	TenantSignature string
	OwnerSignature  string
	Rejected        bool `gorm:"type:bool;default:false"`
	Accepted        bool `gorm:"type:bool;default:false"`
	Signed          bool `gorm:"default:false"`
	Drafted         bool `gorm:"type:bool;default:false"`
}

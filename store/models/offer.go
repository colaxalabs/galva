package models

import (
	"time"
)

type Offer struct {
	ID             *ID       `gorm:"primary_key;not_null"`
	Purpose        string    `gorm:"text;not_null"`
	Size           string    `gorm:"not_null"`
	Duration       time.Time `gorm:"not_null"`
	Cost           string    `gorm:"not_null"`
	Owner          string    `gorm:"not_null"`
	User           *User
	UserAddress    string `gorm:"not_null"`
	Title          string `gorm:"not_null"`
	FullFilled     bool   `gorm:"default:false;not_null"`
	Property       *Property
	PropertyID     int `gorm:"type:integer;not_null"`
	UserSignature  string
	OwnerSignature string
	ExpiresIn      time.Time
	Accepted       bool `gorm:"type:bool;default:false"`
	Signed         bool `gorm:"type:bool;default:false"`
	Drafted        bool `gorm:"type:bool;default:false"`
}

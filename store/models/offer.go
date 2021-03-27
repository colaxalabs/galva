package models

import (
	"time"
)

type Offer struct {
	ID             *ID       `gorm:"primary_key;not_null"`
	Purpose        string    `gorm:"text;not_null"`
	Size           *Wei      `gorm:"type:text"`
	Duration       time.Time `gorm:"not_null"`
	Cost           *Wei      `gorm:"type:text"`
	Owner          string    `gorm:"not_null"`
	User           *User
	UserAddress    string `gorm:"not_null"`
	FullFilled     bool   `gorm:"default:false;not_null"`
	Property       *Property
	PropertyID     int64 `gorm:"type:integer;not_null"`
	UserSignature  string
	OwnerSignature string
	ExpiresIn      time.Time
	Accepted       bool `gorm:"type:bool;default:false"`
	Signed         bool `gorm:"type:bool;default:false"`
	Drafted        bool `gorm:"type:bool;default:false"`
}

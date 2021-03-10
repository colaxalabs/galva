package models

import (
	"time"
)

// Land represent land asset
type Land struct {
	ID            *ID       `gorm:"primary_key;not_null;"`
	TokenId       int       `gorm:"type:integer;not_null;"`
	PostalCode    int       `gorm:"type:integer;not_null;"`
	SateliteImage string    `gorm:"type:varchar(100);not_null;"`
	Location      string    `gorm:"type:varchar(100);not_null;"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP;"`
	UpdatedAt     time.Time
}

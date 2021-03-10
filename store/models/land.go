package models

import (
	"math/big"
	"time"
)

// Land represent land asset
type Land struct {
	ID         *ID       `gorm:"primary_key;not_null;"`
	TokenId    int       `gorm:"type:integer;"`
	Title      string    `gorm:"type:varchar(50);"`
	Size       *big.Int  `gorm:"type:bigint;"`
	SizeUnit   string    `gorm:"type:varchar(10);"`
	PostalCode int       `gorm:"type:integer;"`
	Location   string    `gorm:"type:varchar(100);"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP;"`
	UpdatedAt  time.Time
}

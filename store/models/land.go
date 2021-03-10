package models

import (
	"math/big"
)

// Land represent land asset
type Land struct {
	Model
	TokenId    int      `gorm:"type:integer;"`
	Title      string   `gorm:"type:varchar(50);"`
	Size       *big.Int `gorm:"type:integer;"`
	SizeUnit   string   `gorm:"type:varchar(10);"`
	PostalCode int      `gorm:"type:integer;"`
	Location   string   `gorm:"type:varchar(100);"`
}

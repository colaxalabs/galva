package models

import (
	"time"
)

// Land represent land asset
type Land struct {
	ID            int       `gorm:"primary_key;not_null;"`
	PostalCode    string    `gorm:"type:varchar(10);not_null;"`
	SateliteImage string    `gorm:"type:varchar(100);not_null;"`
	State         string    `gorm:"type:varchar(50);not_null;"`
	Location      string    `gorm:"type:varchar(100);not_null;"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP;"`
	UpdatedAt     time.Time
}

package models

import (
	"time"
)

type User struct {
	ID         *ID       `gorm:"not_null"`
	Address    string    `gorm:"primary_key;type:varchar(50)"`
	Signature  string    `gorm:"type:varchar(225);not_null"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time
	Properties []*Land
}

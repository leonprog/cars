package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model

	Mark       string         `json:"mark" `
	ModelCar   string         `json:"model"`
	OwnerCount int            `json:"owner_count"`
	Price      int            `json:"price"`
	Currency   string         `json:"currency"`
	Options    datatypes.JSON `json:"options" gorm:"type:json"`
}

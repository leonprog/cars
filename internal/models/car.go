package models

import (
	"gorm.io/datatypes"
)

type Car struct {
	Id         int            `json:"id"`
	Mark       string         `json:"mark" `
	ModelCar   string         `json:"model"`
	OwnerCount int            `json:"owner_count"`
	Price      int            `json:"price"`
	Currency   string         `json:"currency"`
	Options    datatypes.JSON `json:"options" gorm:"type:json"`
}

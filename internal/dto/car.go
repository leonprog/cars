package dto

import "gorm.io/datatypes"

type CarCreateDto struct {
	Mark       string
	Model      string
	OwnerCount int
	Price      int
	Currency   string
	Options    datatypes.JSON
}

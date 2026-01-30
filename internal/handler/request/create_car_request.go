package request

import "gorm.io/datatypes"

type CreateCarRequest struct {
	Mark       string         `json:"mark" validate:"min=1"`
	Model      string         `json:"model" validate:"min=1"`
	OwnerCount int            `json:"owner_count" validate:"gt=0"`
	Price      int            `json:"price" validate:"gt=0"`
	Currency   string         `json:"currency" validate:"oneof=RUB USD EUR"`
	Options    datatypes.JSON `json:"options" validate:"required"`
}

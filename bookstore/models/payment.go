package models

import (
	utils "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/utils"
)

type Payment struct {
	Model
	BookID    int      `json:"book_id" validate:"required"`
	Book      Book   `json:"-" gorm:"foreignKey:BookID;constraint:OnDelete:SET NULL;" validate:"-"`
	Quantity  int      `json:"quantity" gorm:"column:quantity;default:1" validate:"required,gt=0"` // default quantity to 1 for simplicity
	// TODO: add user id
}


func (p *Payment) ValidatePayment() error {
	return utils.Validate.Struct(p)
}
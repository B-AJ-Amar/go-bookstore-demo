package models

import (
	"github.com/go-playground/validator/v10"
)

type Book struct {
	Model
	AuthorID      int
	Author      Author   `gorm:"foreignKey:AuthorID";constraint:OnDelete:SET NULL;"`
	Name        string   `gorm:"column:name" validate:"required,max=100,matches=^[a-zA-Z0-9_ ]*$"` 
	Description *string  `gorm:"column:description" validate:"max=500"` 
	Price       float64  `gorm:"column:price" validate:"required,gt=0"` 

}


func (b *Book) ValidateBook() error {
	validate := validator.New()
	return validate.Struct(b)
}
package models

import (
	utils "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/utils"
)

type Book struct {
	Model
	AuthorID    int      `json:"author_id" validate:"required"`
	Author      Author   `json:"-" gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL;" validate:"-"`
	Name        string   `json:"name" gorm:"column:name" validate:"required,alpha_space"`
	Description *string  `json:"description" gorm:"column:description" validate:"max=500"`
	Price       float64  `json:"price" gorm:"column:price" validate:"required,gt=0"`
}


func (b *Book) ValidateBook() error {
	return utils.Validate.Struct(b)
}
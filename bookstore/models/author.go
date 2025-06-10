package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)


type Author struct {
	Model
	Name        string  `gorm:"column:name" validate:"required,maches=^[a-zA-Z ]+$,max=100"` // Name must be alphabetic and max 100 characters
	Email       string  `gorm:"column:email;unique" validate:"email,max=100"` // Email must be unique, valid email format, and max 100 characters
	BirthDate   time.Time `gorm:"column:birth_date"`
	Biography   *string `gorm:"column:biography"`
	Books       []Book  `gorm:"foreignKey:AuthorID"`
}


func (a *Author) ValidateAuthor() error {
	validate := validator.New()
	return validate.Struct(a)
}

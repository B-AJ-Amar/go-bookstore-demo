package models

import (
	"time"

	utils "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/utils"
)


type Author struct {
	Model
	Name        string     `json:"name" gorm:"column:name" validate:"required,alpha_space"` // Name must be alphabetic and max 100 characters
	Email       string     `json:"email" gorm:"column:email;unique" validate:"email,max=100"` // Email must be unique, valid email format, and max 100 characters
	BirthDate   time.Time  `json:"birth_date" gorm:"column:birth_date"`
	Biography   *string    `json:"biography" gorm:"column:biography"`
	Books       []Book     `json:"books" gorm:"foreignKey:AuthorID"`
}


func (a *Author) ValidateAuthor() error {
	return utils.Validate.Struct(a)
}

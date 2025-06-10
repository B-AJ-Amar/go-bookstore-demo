package models


type Book struct {
	Model
	Author      Author	   `gorm:"foreignKey:AuthorID;references:ID"`
	Name        string     `gorm:"column:name"`
	Description *string    `gorm:"column:description"`
	Price       float64    `gorm:"column:price"`

}
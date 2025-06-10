package models

import "time"


type Author struct {
	Model
	Name        string  `gorm:"column:name"`
	BirthDate   time.Time `gorm:"column:birth_date"`
	Biography   *string `gorm:"column:biography"`
}
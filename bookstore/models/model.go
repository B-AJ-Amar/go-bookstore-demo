package models

import (
	"time"

	"gorm.io/gorm"
)

// <-:false : means allow read only, no write
type Model struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;<-:false"`
	CreatedAt time.Time `gorm:"autoCreateTime;<-:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsActive  bool      `gorm:"default:true"` 
}

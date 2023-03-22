package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"not null;uniqueIndex;primaryKey" json:"id"`
	Title       string         `gorm:"size: 255;not null;" json:"title"`
	Description string         `gorm:"null;" json:"description"`
	Author      string         `gorm:"size: 100;null;" json:"author"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

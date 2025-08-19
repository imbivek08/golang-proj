package models

import "time"

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"not null"`
	CreatedAt time.Time
}

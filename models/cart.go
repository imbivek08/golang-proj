package models

import "time"

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	CreatedAt time.Time
	Items     []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Quantity  uint `gorm:"not null"`
}

package models

import "time"

type Order struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
	CreatedAt  time.Time
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Subtotal  float64 `gorm:"not null"`
}

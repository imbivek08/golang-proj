package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FullName  string
	Address   string
	Phone     string
	Role      string `gorm:"default:'user'"`
	CreatedAt time.Time
	Orders    []Order `gorm:"foreignKey:UserID"`
}

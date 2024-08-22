package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Balance   int    `gorm:"default:0"`
	CreatedAt time.Time
}

package models

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:255;not null"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `json:"user_id" gorm:"not null"`
	CreatedAt time.Time
	User      User `json:"-" gorm:"foreignKey:UserID"`
}

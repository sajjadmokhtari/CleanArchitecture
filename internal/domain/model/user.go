package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Phone     string    `gorm:"uniqueIndex;size:11"`
	CreatedAt time.Time
}

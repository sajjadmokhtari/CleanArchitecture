package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone string `gorm:"uniqueIndex;size:11"`
	Role  string
}

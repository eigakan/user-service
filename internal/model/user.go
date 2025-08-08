package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Login    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Logo     string
	Settings UserSettings `gorm:"type:json"`
}

type UserSettings struct {
	NotificationsEnabled bool
}

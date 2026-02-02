package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	FirstName string
	LastName string
	Phone *string
	BirthDate *time.Time
	IsActive bool
}
package models

import (
	"time"

	"gorm.io/gorm"
)

type IsActive bool

const (
	Active    IsActive = true
	NotActive IsActive = false
)

type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"type:varchar(100)"`
	LastName  string `gorm:"type:varchar(100)"`
	Phone     *string
	BirthDate *time.Time
	IsActive  bool `gorm:"default:Active"`
}

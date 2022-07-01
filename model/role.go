package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"size:50;not null"`
	DisplayName string `gorm:"size:128;not null"`
}

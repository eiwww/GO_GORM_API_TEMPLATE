package model

import (
	"gorm.io/gorm"
)

type Province struct {
	gorm.Model
	Name     string     `gorm:"size:45;not null"`
	Section  sectorType `gorm:"type:ENUM('center', 'north', 'south');not null"`
}

package models

import (
	"gorm.io/gorm"
)

type Province struct {
	gorm.Model
	Name     string     `gorm:"size:45;not null"`
	Section  sectorType `gorm:"type:ENUM('center', 'north', 'south');not null"`
	IsDelete isDelete   `gorm:"type:ENUM('no', 'yes');not null;default:'no'"`
}

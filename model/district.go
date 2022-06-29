package model

import (
	"gorm.io/gorm"
)

type District struct {
	gorm.Model
	Province   Province
	ProvinceId uint
	Name       string   `gorm:"size:45;not null"`
	IsDelete   isDelete `gorm:"type:ENUM('no', 'yes');not null;default:'no'"`
}

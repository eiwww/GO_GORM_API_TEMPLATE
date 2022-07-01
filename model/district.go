package model

import (
	"gorm.io/gorm"
)

type District struct {
	gorm.Model
	Province   Province
	ProvinceId uint
	Name       string `gorm:"size:45;not null"`
}

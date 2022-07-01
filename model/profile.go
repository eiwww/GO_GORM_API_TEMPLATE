package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserId     string     `gorm:"size:191;not null"`
	Firstname  string     `gorm:"size:100;not null"`
	Lastname   string     `gorm:"size:100;not null"`
	Gender     genderType `gorm:"not null"`
	Phone      string     `gorm:"size:50;not null"`
	Province   Province
	ProvinceId uint
	District   District
	DistrictId uint
	Village    string         `gorm:"size:100;"`
	Status     userStatusType `gorm:"not null"`
}

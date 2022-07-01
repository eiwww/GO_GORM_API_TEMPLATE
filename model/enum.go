package model

import "database/sql/driver"

type sectorType string

const (
	center sectorType = "center"
	north  sectorType = "north"
	south  sectorType = "south"
)

func (st *sectorType) Scan(value interface{}) error {
	*st = sectorType(value.([]byte))
	return nil
}

func (st sectorType) Value() (driver.Value, error) {
	return string(st), nil
}

type genderType string

const (
	m genderType = "m"
	f genderType = "f"
)

func (gt *genderType) Scan(value interface{}) error {
	*gt = genderType(value.([]byte))
	return nil
}

func (gt genderType) Value() (driver.Value, error) {
	return string(gt), nil
}

type userStatusType string

const (
	active  userStatusType = "active"
	offline userStatusType = "offline"
)

func (ust *userStatusType) Scan(value interface{}) error {
	*ust = userStatusType(value.([]byte))
	return nil
}

func (ust userStatusType) Value() (driver.Value, error) {
	return string(ust), nil
}

// type isDelete string

// const (
// 	no  isDelete = "no"
// 	yes isDelete = "no"
// )

// func (de *isDelete) Scan(value interface{}) error {
// 	*de = isDelete(value.([]byte))
// 	return nil
// }

// func (de isDelete) Value() (driver.Value, error) {
// 	return string(de), nil
// }

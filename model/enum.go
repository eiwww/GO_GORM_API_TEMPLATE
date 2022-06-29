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

type isDelete string

const (
	no  isDelete = "no"
	yes isDelete = "no"
)

func (de *isDelete) Scan(value interface{}) error {
	*de = isDelete(value.([]byte))
	return nil
}

func (de isDelete) Value() (driver.Value, error) {
	return string(de), nil
}

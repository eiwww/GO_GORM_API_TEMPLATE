package seeder

import (
	"fmt"
	"rest/database"
	model "rest/model"
)

type Province = model.Province

type District = model.District

type Role = model.Role

func SeedDB() {

	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		return
	}

	province := []model.Province{}
	proData := db.Find(&province)
	if proData.RowsAffected <= 0 {
		db.Create(&provinces)
	}

	district := []model.District{}
	disData := db.Find(&district)
	if disData.RowsAffected <= 0 {
		db.Create(&districts)
	}

	role := []model.Role{}
	roleData := db.Find(&role)
	if roleData.RowsAffected <= 0 {
		db.Create(&roles)
	}
}

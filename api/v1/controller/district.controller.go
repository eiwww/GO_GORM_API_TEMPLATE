package controllerV1

import (
	"errors"
	"fmt"
	"rest/database"
	model "rest/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type District = model.District

func GetAllDistrict(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	provinceId := c.Param("provinceId")

	districts := []model.District{}
	tx := db.Joins("Province").Where("province_id = ?", provinceId).Find(&districts)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	c.JSON(200, gin.H{"districts": districts})
}

func GetDistrictById(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	districtId := c.Param("districtId")

	districts := model.District{}
	tx := db.Joins("Province").Where("`districts`.`id` = ?", districtId).Find(&districts)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	if tx.RowsAffected <= 0 {
		c.JSON(200, "")
		return
	}

	c.JSON(200, districts)
}

func CreateDistrict(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	tx := db.Begin()

	body := District{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	district := District{
		ProvinceId: body.ProvinceId,
		Name:       body.Name,
	}

	if err := tx.Create(&district).Error; err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(201, gin.H{"message": "Insert Data Complete"})
}

func UpdateDistrict(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	tx := db.Begin()

	body := District{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var disModel District

	disRes := tx.First(&disModel, "id = ?", body.ID)

	if disRes.Error != nil {
		fmt.Println(disRes.Error)
		if errors.Is(disRes.Error, gorm.ErrRecordNotFound) {
			c.JSON(400, gin.H{
				"message": "Record not found",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Error occurred while updating user",
			})
		}
		return
	}

	disModel.Name = body.Name
	disModel.ProvinceId = body.ProvinceId

	if err := tx.Save(&disModel).Error; err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	c.JSON(200, gin.H{"message": "Update data complete"})
}

func DeleteDistrict(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	tx := db.Begin()

	districtId := c.Param("districtId")

	var disModel District

	disRes := tx.First(&disModel, "id = ?", districtId)

	if disRes.Error != nil {
		fmt.Println(disRes.Error)
		if errors.Is(disRes.Error, gorm.ErrRecordNotFound) {
			c.JSON(400, gin.H{
				"message": "Record not found",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Error occurred while updating user",
			})
		}
		return
	}

	if err := tx.Delete(&disModel).Error; err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	c.JSON(200, gin.H{"message": "Delete data complete"})
}

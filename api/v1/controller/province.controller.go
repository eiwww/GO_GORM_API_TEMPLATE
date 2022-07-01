package controllerV1

import (
	"errors"
	"fmt"
	"rest/database"
	model "rest/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Province = model.Province

func GetAllProvince(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	provinces := []Province{}
	tx := db.Find(&provinces)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	c.JSON(200, gin.H{"provinces": provinces})
}

func GetProvinceById(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	provinceId := c.Param("provinceId")

	provinces := Province{}
	tx := db.Where("id = ?", provinceId).Find(&provinces)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	c.JSON(200, gin.H{"provinces": provinces})
}

func CreateProvince(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	tx := db.Begin()

	body := Province{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	province := Province{
		Name:    body.Name,
		Section: body.Section,
	}

	if err := tx.Create(&province).Error; err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(201, gin.H{"message": "Insert Data Complete"})
}

func UpdateProvince(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	tx := db.Begin()

	body := Province{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var proModel Province

	proRes := tx.First(&proModel, "id = ?", body.ID)

	if proRes.Error != nil {
		fmt.Println(proRes.Error)
		if errors.Is(proRes.Error, gorm.ErrRecordNotFound) {
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

	proModel.Name = body.Name
	proModel.Section = body.Section

	if err := tx.Save(&proModel).Error; err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	c.JSON(200, gin.H{"message": "Update data complete"})
}

func DeleteProvince(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	tx := db.Begin()

	provinceId := c.Param("provinceId")

	var proModel Province

	proRes := tx.First(&proModel, "id = ?", provinceId)

	if proRes.Error != nil {
		fmt.Println(proRes.Error)
		if errors.Is(proRes.Error, gorm.ErrRecordNotFound) {
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

	if err := tx.Delete(&proModel).Error; err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	c.JSON(200, gin.H{"message": "Delete data complete"})
}

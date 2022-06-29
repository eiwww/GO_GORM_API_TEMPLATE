package controllerV1

import (
	"fmt"
	"rest/database"
	model "rest/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

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
	tx := db.Preload(clause.Associations).Where("province_id = ?", provinceId).Find(&districts)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	c.JSON(200, gin.H{"districts": districts})
}

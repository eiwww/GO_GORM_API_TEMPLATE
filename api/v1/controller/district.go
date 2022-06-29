package controllerV1

import (
	"fmt"
	"rest/database"
	model "rest/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type District struct {
	ID         uint      `json:"districtId"`
	ProvinceID uint      `json:"provinceId"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

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
	tx := db.Preload(clause.Associations).Where("id = ?", districtId).Find(&districts)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	c.JSON(200, gin.H{"districts": districts})
}

func CreateDistrict(c *gin.Context) {
	// db, conErr := database.GetDatabaseConnection()
	// if conErr != nil {
	// 	fmt.Println(conErr)
	// 	c.JSON(400, gin.H{
	// 		"message": "Service is unavailable",
	// 	})
	// 	return
	// }

	json := District{}
	c.BindJSON(&json)

	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(400, gin.H{
	// 		"message": "No Request Body Data",
	// 	})
	// 	return
	// }

	fmt.Println(json)

	// districts := model.District{}
	// tx := db.Preload(clause.Associations).Where("id = ?", districtId).Find(&districts)
	// if tx.Error != nil {
	// 	fmt.Println(tx.Error)
	// 	return
	// }
	c.JSON(200, gin.H{"message": json})
}

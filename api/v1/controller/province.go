package controllerV1

import (
	"fmt"
	"rest/database"
	model "rest/model"

	"github.com/gin-gonic/gin"
)

func GetAllProvince(c *gin.Context) {
	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		c.JSON(400, gin.H{
			"message": "Service is unavailable",
		})
		return
	}

	provinces := []model.Province{}

	tx := db.Find(&provinces)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	c.JSON(200, gin.H{"provinces": provinces})
}

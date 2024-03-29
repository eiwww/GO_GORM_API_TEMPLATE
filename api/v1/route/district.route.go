package route

import (
	controllerV1 "rest/api/v1/controller"

	"github.com/gin-gonic/gin"
)

// Function to setup routers and router groups
func DistrictRoutes(app *gin.RouterGroup) {

	district := app.Group("/districts")
	{
		district.GET("/:provinceId", controllerV1.GetAllDistrict)
		district.GET("/get/:districtId", controllerV1.GetDistrictById)
		district.POST("/", controllerV1.CreateDistrict)
		district.PUT("/", controllerV1.UpdateDistrict)
		district.POST("/delete/:districtId", controllerV1.DeleteDistrict)
	}

}

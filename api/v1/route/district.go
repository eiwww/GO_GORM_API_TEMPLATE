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
	}

}
package route

import (
	controllerV1 "rest/api/v1/controller"

	"github.com/gin-gonic/gin"
)

// Function to setup routers and router groups
func ProvinceRoutes(app *gin.RouterGroup) {

	province := app.Group("/provinces")
	{
		province.GET("/", controllerV1.GetAllProvince)
	}

}

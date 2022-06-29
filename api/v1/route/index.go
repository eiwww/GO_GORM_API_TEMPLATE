package route

import "github.com/gin-gonic/gin"

func AddRoutes(route *gin.RouterGroup) {
	ProvinceRoutes(route)
	DistrictRoutes(route)
}

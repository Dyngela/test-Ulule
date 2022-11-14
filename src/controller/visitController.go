package controller

import (
	"Ulule/src/service"
	"github.com/gin-gonic/gin"
)

// VisitController Every API routes mainly related to Visit table
func VisitController(router *gin.Engine) {
	v1 := router.Group("/api/v1/visit/")
	{
		v1.GET("visits/project/:id/date/:date/range/:range", service.GetVisitsByProjectByDate)
	}
}

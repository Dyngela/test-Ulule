package controller

import (
	"Ulule/src/service"
	"github.com/gin-gonic/gin"
)

// ProjectController Every API routes mainly related to Project table
func ProjectController(router *gin.Engine) {
	v1 := router.Group("/api/v1/project/")
	{
		v1.GET("milestones/project/:id", service.GetProjectMilestones)
		v1.GET("advancement/percentage/project/:id/date/:date/range/:range", service.GetAdvancementPercentage)
	}
}

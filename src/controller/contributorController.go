package controller

import (
	"Ulule/src/service"
	"github.com/gin-gonic/gin"
)

func ContributionController(router *gin.Engine) {
	v1 := router.Group("/api/v1/contribution/")
	{
		v1.GET("amount/project/:id/date/:date/range/:range", service.GetContributionByDateByProject)
		v1.GET("contributors/project/:id/date/:date/range/:range", service.GetNewContributorByDateByProject)
		v1.GET("contributions/project/:id/date/:date/range/:range", service.GetNewContributionByDateByProject)
		v1.GET("average/contribution/project/:id", service.GetAverageContributionAmount)
		v1.GET("rate/visistor/project/:id/date/:date/range/:range", service.GetContributionRateByVisitorsByDateByProject)
	}
}

package service

import (
	"Ulule/src/exception"
	"Ulule/src/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetProjectMilestones(c *gin.Context) {
	type URI struct {
		ProjectId int `binding:"gte=1" uri:"id"`
	}
	uri := URI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		exception.ThrowExceptionBadArgument(c, err)
		return
	}

	response, err := repository.FindProjectMilestones(uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetAdvancementPercentage(c *gin.Context) {
	type URI struct {
		Date      time.Time `binding:"required" uri:"date" time_format:"2006-01-02"`
		ProjectId int       `binding:"gte=1" uri:"id"`
		DateRange int       `binding:"gte=0" uri:"range"`
	}
	uri := URI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		exception.ThrowExceptionBadArgument(c, err)
		return
	}

	response, err := repository.FindAdvancementPercentage(
		uri.Date,
		uri.DateRange,
		uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
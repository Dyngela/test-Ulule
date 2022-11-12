package service

import (
	"Ulule/src/exception"
	"Ulule/src/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func GetContributionByDateByProject(c *gin.Context) {
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

	response, err := repository.FindContributionByDateByProject(
		strings.Join(getDateRangeInString(uri.Date, uri.DateRange), "', '"),
		uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetNewContributorByDateByProject(c *gin.Context) {
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

	response, err := repository.FindNewContributorByDateByProject(
		getDateRangeInString(uri.Date, uri.DateRange),
		uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetNewContributionByDateByProject(c *gin.Context) {
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

	response, err := repository.FindNewContributionByDateByProject(
		strings.Join(getDateRangeInString(uri.Date, uri.DateRange), "', '"),
		uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func getDateRangeInString(date time.Time, dateRange int) []string {
	var listOfDate []string

	for i := 0; i < dateRange+1; i++ {
		tempDate := date.AddDate(0, 0, i*-1)
		listOfDate = append(listOfDate, tempDate.Format("2006-01-02"))
	}
	return listOfDate
}

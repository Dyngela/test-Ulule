package service

import (
	"Ulule/src/exception"
	"Ulule/src/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetContributionByDateByProject Service to get the sum of the contributions each day on a range of date for a given project
// @Summary Get each day specified and send back the sum of the contributions for that day
// @Description Give an array containing the percentage of advancement for a given day
// @Tags Contribution
// @Param id path int true "Project's ID"
// @Param date path string true "Last date of the response with the format YYYY-MM-DD"
// @Param range path int true "On how many day you want before the given date"
// @Success 200 {array} DTO.ContributionByDateByProject
// @Failure 500 {array} object
// @Failure 400 {array} object
// @Router /contribution/amount/project/{id}/date/{date}/range/{range} [get]
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
		uri.Date,
		uri.DateRange,
		uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetNewContributorByDateByProject Service to get the count of new contributors each day on a range of date for a given project
// @Summary Get each day specified and send back the count of new contributor for that day
// @Description Give an array containing the count of new contributor for a given day
// @Tags Contribution
// @Param id path int true "Project's ID"
// @Param date path string true "Last date of the response with the format YYYY-MM-DD"
// @Param range path int true "On how many day you want before the given date"
// @Success 200 {array} DTO.NewContributorByDateByProject
// @Failure 500 {array} object
// @Failure 400 {array} object
// @Router /contribution/contributors/project/{id}/date/{date}/range/{range} [get]
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
		uri.Date,
		uri.DateRange,
		uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetNewContributionByDateByProject Service to get the count of new contributions each day on a range of date for a given project
// @Summary Get each day specified and send back the count of new contributions for that day
// @Description Give an array containing the count of new contribution for a given day
// @Tags Contribution
// @Param id path int true "Project's ID"
// @Param date path string true "Last date of the response with the format YYYY-MM-DD"
// @Param range path int true "On how many day you want before the given date"
// @Success 200 {array} DTO.NewContributionByDateByProject
// @Failure 500 {array} object
// @Failure 400 {array} object
// @Router /contribution/contributions/project/{id}/date/{date}/range/{range} [get]
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
		uri.Date,
		uri.DateRange,
		uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetAverageContributionAmount Service to get the average of the contribution for a given project
// @Summary Get the average of the contribution for a given project
// @Description Get the average of the contribution for a given project
// @Tags Contribution
// @Param id path int true "Project's ID"
// @Param date path string true "Last date of the response with the format YYYY-MM-DD"
// @Param range path int true "On how many day you want before the given date"
// @Success 200 {object} DTO.AverageContributionByProject
// @Failure 500 {object} object
// @Failure 400 {object} object
// @Router /contribution/average/contribution/project/{id} [get]
func GetAverageContributionAmount(c *gin.Context) {
	type URI struct {
		ProjectId int `binding:"gte=1" uri:"id"`
	}
	uri := URI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		exception.ThrowExceptionBadArgument(c, err)
		return
	}

	response, err := repository.FindAverageContributionAmount(uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetContributionRateByVisitorsByDateByProject Service to get the rate of contribution according to the number of visitors
// @Summary Get each day specified and send back the rate between visitor and contribution
// @Description Give an array containing the contribution rate for a given day and a given project
// @Tags Contribution
// @Param id path int true "Project's ID"
// @Param date path string true "Last date of the response with the format YYYY-MM-DD"
// @Param range path int true "On how many day you want before the given date"
// @Success 200 {array} DTO.ContributionRateByVisitorsByDateByProject
// @Failure 500 {array} object
// @Failure 400 {array} object
// @Router /contribution/rate/visitor/project/{id}/date/{date}/range/{range} [get]
func GetContributionRateByVisitorsByDateByProject(c *gin.Context) {
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

	response, err := repository.FindContributionRateByVisitorsByDateByProject(
		uri.Date,
		uri.DateRange,
		uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

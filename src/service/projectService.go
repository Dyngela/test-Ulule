package service

import (
	"Ulule/src/exception"
	"Ulule/src/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetProjectMilestones Service to get the milestone of a project
// @Summary Get each milestone from a project, since first contribution to 100% completion
// @Description Give the contribution ID, the date and the slice completed of a project objective.
// @Tags Project
// @Param id path int true "Project's ID"
// @Success 200 {array} DTO.ProjectMilestone
// @Failure 500 {array} object
// @Failure 400 {array} object
// @Router /project/milestones/project/{id} [get]
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

// GetAdvancementPercentage Service to get the advancement of the project in percentage each day
// @Summary Get each day specified and send back the percentage reach each day.
// @Description Give an array containing the percentage of advancement for a given day
// @Tags Project
// @Param id path int true "Project's ID"
// @Param date path string true "Last date of the response with the format YYYY-MM-DD"
// @Param range path int true "On how many day you want before the given date"
// @Success 200 {array} DTO.ProjectMilestone
// @Failure 500 {array} object
// @Failure 400 {array} object
// @Router /advancement/percentage/project/{id}/date/{date}/range/{range} [get]
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

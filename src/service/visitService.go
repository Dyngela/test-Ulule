package service

import (
	"Ulule/src/exception"
	"Ulule/src/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetVisitsByProjectByDate Service to get the number of visitors.
// @Summary Get the number of visitor
// @Description Get the number of visitor on a range of date for a given project
// @Tags Visit
// @Param id path int true "Project's ID"
// @Param date path string true "Last date of the response with the format YYYY-MM-DD"
// @Param range path int true "On how many day you want before the given date"
// @Success 200 {array} DTO.VisitByDateByProject
// @Failure 500 {array} object
// @Failure 400 {array} object
// @Router /visit/visits/project/{id}/date/{date}/range/{range} [get]
func GetVisitsByProjectByDate(c *gin.Context) {
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

	response, err := repository.FindNumberOfVisitsByDateByProject(
		uri.Date,
		uri.DateRange,
		uri.ProjectId)
	if err != nil {
		exception.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

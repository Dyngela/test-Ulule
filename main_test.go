package main

import (
	"Ulule/src/DTO"
	"Ulule/src/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetVisitsByProjectByDateHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/visit/visits/project/:id/date/:date/range/:range", service.GetVisitsByProjectByDate)
	req, _ := http.NewRequest("GET", "/api/v1/visit/visits/project/44861/date/2022-02-19/range/5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.VisitByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)
	date, err := time.Parse("2006-01-02", "2022-02-19")
	if err != nil {
		t.Failed()
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, 5, len(resp))
	assert.GreaterOrEqual(t, date, resp[0].Date.Time)
	assert.NotEmpty(t, resp)
}

func TestGetVisitsByProjectByDateHandlerFail(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/visit/visits/project/:id/date/:date/range/:range", service.GetVisitsByProjectByDate)
	req, _ := http.NewRequest("GET", "/api/v1/visit/visits/project/0/date/20-12-12/range/0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.VisitByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, resp, []DTO.VisitByDateByProject(nil))
}

func TestGetProjectMilestonesSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/project/milestones/project/:id", service.GetProjectMilestones)
	req, _ := http.NewRequest("GET", "/api/v1/project/milestones/project/44861", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.ProjectMilestone
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, 10, len(resp))
	assert.NotEmpty(t, resp)
}

func TestGetProjectMilestonesFail(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/project/milestones/project/:id", service.GetProjectMilestones)
	req, _ := http.NewRequest("GET", "/api/v1/project/milestones/project/0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.ProjectMilestone
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, resp, []DTO.ProjectMilestone(nil))
}

func TestGetAdvancementPercentageSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/advancement/percentage/project/:id/date/:date/range/:range", service.GetAdvancementPercentage)
	req, _ := http.NewRequest("GET", "/api/v1/advancement/percentage/project/44861/date/2017-06-10/range/10", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	date, err := time.Parse("2006-01-02", "2022-02-19")
	if err != nil {
		t.Failed()
	}

	var resp []DTO.ProjectAdvancement
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, 10, len(resp))
	assert.GreaterOrEqual(t, date, resp[0].Date.Time)
	assert.NotEmpty(t, resp)
}

func TestGetAdvancementPercentageFail(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/advancement/percentage/project/:id/date/:date/range/:range", service.GetAdvancementPercentage)
	req, _ := http.NewRequest("GET", "/api/v1/advancement/percentage/project/44861/date/01-12-12/range/86", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.ProjectAdvancement
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, resp, []DTO.ProjectAdvancement(nil))
}

func TestGetContributionByDateByProjectSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/amount/project/:id/date/:date/range/:range", service.GetContributionByDateByProject)
	req, _ := http.NewRequest("GET", "/api/v1/amount/project/44861/date/2017-06-10/range/10", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.ContributionByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, 10, len(resp))
	assert.NotEmpty(t, resp)
}

func TestGetContributionByDateByProjectFail(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/amount/project/:id/date/:date/range/:range", service.GetAdvancementPercentage)
	req, _ := http.NewRequest("GET", "/api/v1/amount/project/44861/date/2017-13-10/range/0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.ContributionByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, resp, []DTO.ContributionByDateByProject(nil))
}

func TestGetNewContributorByDateByProjectSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/contributors/project/:id/date/:date/range/:range", service.GetNewContributorByDateByProject)
	req, _ := http.NewRequest("GET", "/api/v1/contributors/project/44861/date/2017-06-10/range/10", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.NewContributorByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, 10, len(resp))
	assert.NotEmpty(t, resp)
}

func TestGetNewContributorByDateByProjectFail(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/contributors/project/:id/date/:date/range/:range", service.GetNewContributorByDateByProject)
	req, _ := http.NewRequest("GET", "/api/v1/contributors/project/0/date/2017-31-10/range/0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.NewContributorByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, resp, []DTO.NewContributorByDateByProject(nil))
}

func TestGetNewContributionByDateByProjectSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/contribution/contributions/project/:id/date/:date/range/:range", service.GetNewContributionByDateByProject)
	req, _ := http.NewRequest("GET", "/api/v1/contribution/contributions/project/44861/date/2017-06-10/range/10", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.NewContributionByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, 10, len(resp))
	assert.NotEmpty(t, resp)
}

func TestGetNewContributionByDateByProjectFail(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/contribution/contributions/project/:id/date/:date/range/:range", service.GetNewContributionByDateByProject)
	req, _ := http.NewRequest("GET", "/api/v1/contribution/contributions/project/0/date/2017-15-10/range/0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.NewContributionByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, resp, []DTO.NewContributionByDateByProject(nil))
}

func TestGetAverageContributionAmountSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/average/contribution/project/:id", service.GetAverageContributionAmount)
	req, _ := http.NewRequest("GET", "/api/v1/average/contribution/project/44861", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp DTO.AverageContributionByProject
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resp)
}

func TestGetAverageContributionAmountFail(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/average/contribution/project/:id", service.GetAverageContributionAmount)
	req, _ := http.NewRequest("GET", "/api/v1/average/contribution/project/0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp DTO.AverageContributionByProject
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Empty(t, resp)
}

func TestGetContributionRateByVisitorsByDateByProjectSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/rate/visitor/project/:id/date/:date/range/:range", service.GetContributionRateByVisitorsByDateByProject)
	req, _ := http.NewRequest("GET", "/api/v1/rate/visitor/project/135561/date/2022-02-22/range/5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.ContributionRateByVisitorsByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, 5, len(resp))
	assert.NotEmpty(t, resp)
}

func TestGetContributionRateByVisitorsByDateByProjectFail(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/rate/visitor/project/:id/date/:date/range/:range", service.GetContributionRateByVisitorsByDateByProject)
	req, _ := http.NewRequest("GET", "/api/v1/rate/visitor/project/0/date/2017-15-12/range/0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []DTO.ContributionRateByVisitorsByDateByProject
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, resp, []DTO.ContributionRateByVisitorsByDateByProject(nil))
}

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
	assert.LessOrEqual(t, len(resp), 5)
	assert.LessOrEqual(t, resp[0].Date.Time, date)
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

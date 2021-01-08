package api

import (
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// FindReport api handler
type FindReport struct {
	reportProvider provider.Report
}

// NewFindReport find previously created report
func NewFindReport(reportProvider provider.Report) *FindReport {
	return &FindReport{reportProvider: reportProvider}
}

// Path return api path
func (r *FindReport) Path() string {
	return "/reports"
}

// Method return api method
func (r *FindReport) Method() string {
	return "GET"
}

// Handle request otp
func (r *FindReport) Handle(context provider.APIContext) {
	userID := context.Get("user-id").(int)
	if userID <= 0 {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	data, err := r.reportProvider.FindByUserID(context, userID)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{"data": data})
}

package api

import (
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportMeDelete api handler
type ReportMeDelete struct {
	reportProvider provider.Report
}

// NewReportMeDelete find previously created report
func NewReportMeDelete(reportProvider provider.Report) *ReportMeDelete {
	return &ReportMeDelete{reportProvider: reportProvider}
}

// Path return api path
func (r *ReportMeDelete) Path() string {
	return "/reports"
}

// Method return api method
func (r *ReportMeDelete) Method() string {
	return "DELETE"
}

// Handle request otp
func (r *ReportMeDelete) Handle(context provider.APIContext) {
	userID := context.Get("user-id").(int)
	if userID <= 0 {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	err := r.reportProvider.DeleteReportedCases(context, userID)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
	}

	_ = context.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Report has been successfully deleted",
	})
}

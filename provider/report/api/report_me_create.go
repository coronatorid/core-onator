package api

import (
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportMeCreate api handler
type ReportMeCreate struct {
	reportProvider provider.Report
}

// NewReportMeCreate create new request otp handler object
func NewReportMeCreate(reportProvider provider.Report) *ReportMeCreate {
	return &ReportMeCreate{reportProvider: reportProvider}
}

// Path return api path
func (r *ReportMeCreate) Path() string {
	return "/reports"
}

// Method return api method
func (r *ReportMeCreate) Method() string {
	return "POST"
}

// Handle request otp
func (r *ReportMeCreate) Handle(context provider.APIContext) {
	userID := context.Get("user-id").(int)
	if userID <= 0 {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	fileHeader, err := context.FormFile("file")
	if err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	if err := r.reportProvider.CreateReportedCases(context, userID, fileHeader); err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Report has been successfully submitted",
	})
}

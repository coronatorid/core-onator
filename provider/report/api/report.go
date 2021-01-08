package api

import (
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Report api handler
type Report struct {
	reportProvider provider.Report
}

// NewReport create new request otp handler object
func NewReport(reportProvider provider.Report) *Report {
	return &Report{reportProvider: reportProvider}
}

// Path return api path
func (r *Report) Path() string {
	return "/reports"
}

// Method return api method
func (r *Report) Method() string {
	return "POST"
}

// Handle request otp
func (r *Report) Handle(context provider.APIContext) {
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

	if err := r.reportProvider.CreateReportAndUploadFile(context, userID, fileHeader); err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
	}

	context.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Report has been successfully submitted",
	})
}

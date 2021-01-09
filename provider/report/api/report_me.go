package api

import (
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportMe api handler
type ReportMe struct {
	reportProvider provider.Report
}

// NewReportMe find previously created report
func NewReportMe(reportProvider provider.Report) *ReportMe {
	return &ReportMe{reportProvider: reportProvider}
}

// Path return api path
func (r *ReportMe) Path() string {
	return "/reports"
}

// Method return api method
func (r *ReportMe) Method() string {
	return "GET"
}

// Handle request otp
func (r *ReportMe) Handle(context provider.APIContext) {
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
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{"data": data})
}

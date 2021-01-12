package api

import (
	"net/http"
	"strconv"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportDelete api handler
type ReportDelete struct {
	adminProvider provider.Admin
}

// NewReportDelete find previously created report
func NewReportDelete(adminProvider provider.Admin) *ReportDelete {
	return &ReportDelete{adminProvider: adminProvider}
}

// Path return api path
func (r *ReportDelete) Path() string {
	return "/administrations/reports/:id"
}

// Method return api method
func (r *ReportDelete) Method() string {
	return "DELETE"
}

// Handle request otp
func (r *ReportDelete) Handle(context provider.APIContext) {
	userID := context.Get("user-id").(int)
	if userID <= 0 {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	if err := r.adminProvider.ReportDelete(context, userID, id); err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Report has been successfully deleted",
	})
}

package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportStatus api handler
type ReportStatus struct {
	adminProvider provider.Admin
}

// NewReportStatus create new report list handler
func NewReportStatus(adminProvider provider.Admin) *ReportStatus {
	return &ReportStatus{adminProvider: adminProvider}
}

// Path return api path
func (r *ReportStatus) Path() string {
	return "/administrations/reports/:id/status"
}

// Method return api method
func (r *ReportStatus) Method() string {
	return "PATCH"
}

// Handle report list
func (r *ReportStatus) Handle(context provider.APIContext) {
	var request entity.ReportStatus
	if err := json.NewDecoder(context.Request().Body).Decode(&request); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

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

	switch request.Status {
	case constant.ReportedCasesConfirmed.Humanized():
		reportedCase, err := r.adminProvider.ReportConfirm(context, userID, id)
		if err != nil {
			_ = context.JSON(err.HTTPStatus, map[string]interface{}{
				"errors":  err.ErrorString(),
				"message": err.Error(),
			})
			return
		}

		_ = context.JSON(http.StatusOK, map[string]interface{}{
			"data": reportedCase,
		})
	case constant.ReportedCasesRejected.Humanized():
		reportedCase, err := r.adminProvider.ReportReject(context, userID, id)
		if err != nil {
			_ = context.JSON(err.HTTPStatus, map[string]interface{}{
				"errors":  err.ErrorString(),
				"message": err.Error(),
			})
			return
		}

		_ = context.JSON(http.StatusOK, map[string]interface{}{
			"data": reportedCase,
		})
	default:
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
	}

}

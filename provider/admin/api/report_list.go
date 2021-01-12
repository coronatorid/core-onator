package api

import (
	"net/http"
	"strconv"

	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportList api handler
type ReportList struct {
	adminProvider provider.Admin
}

// NewReportList create new report list handler
func NewReportList(adminProvider provider.Admin) *ReportList {
	return &ReportList{adminProvider: adminProvider}
}

// Path return api path
func (r *ReportList) Path() string {
	return "/administrations/reports"
}

// Method return api method
func (r *ReportList) Method() string {
	return "GET"
}

// Handle report list
func (r *ReportList) Handle(context provider.APIContext) {
	userID := context.Get("user-id").(int)
	if userID <= 0 {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	var status string
	var requestMeta entity.RequestMeta

	limitString := context.QueryParam("limit")
	if len(limitString) == 0 {
		requestMeta.Limit = 10
	} else {
		limit, err := strconv.Atoi(limitString)
		if err != nil {
			_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
				"errors":  []entity.APIError{entity.ErrorBadRequest()},
				"message": "Bad request",
			})
			return
		}

		if limit > 10 {
			requestMeta.Limit = 10
		} else {
			requestMeta.Limit = limit
		}
	}

	offsetString := context.QueryParam("offset")
	if len(offsetString) == 0 {
		requestMeta.Offset = 0
	} else {
		offset, err := strconv.Atoi(offsetString)
		if err != nil {
			_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
				"errors":  []entity.APIError{entity.ErrorBadRequest()},
				"message": "Bad request",
			})
			return
		}

		requestMeta.Offset = offset
	}

	status = context.QueryParam("status")
	convertedStatus := constant.ReportedCasesPending
	switch status {
	case "":
	case constant.ReportedCasesConfirmed.Humanized():
		convertedStatus = constant.ReportedCasesConfirmed
	case constant.ReportedCasesRejected.Humanized():
		convertedStatus = constant.ReportedCasesRejected
	case constant.ReportedCasesPending.Humanized():
		convertedStatus = constant.ReportedCasesPending
	default:
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	reportedCases, responseMeta, err := r.adminProvider.ReportList(context, userID, convertedStatus, requestMeta)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{"data": reportedCases, "meta": responseMeta})
}

package usecase

import (
	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportList of submitted report
type ReportList struct{}

// Perform report list logic
func (r *ReportList) Perform(ctx provider.Context, userID int, status constant.ReportedCasesStatus, requestMeta entity.RequestMeta, adminProvider provider.Admin, reportProvider provider.Report) ([]entity.ReportedCases, entity.ResponseMeta, *entity.ApplicationError) {
	var reportedCases []entity.ReportedCases
	var err *entity.ApplicationError
	var total int

	var responseMeta entity.ResponseMeta
	responseMeta.Limit = requestMeta.Limit
	responseMeta.Offset = requestMeta.Offset

	if _, err = adminProvider.Authenticate(ctx, userID, []constant.UserRole{constant.UserRoleSuperAdmin}); err != nil {
		return reportedCases, responseMeta, err
	}

	reportedCases, err = reportProvider.List(ctx, status, requestMeta)
	if err != nil {
		return reportedCases, responseMeta, err
	}

	total, err = reportProvider.Count(ctx, status)
	if err != nil {
		return reportedCases, responseMeta, err
	}
	responseMeta.Total = total
	return reportedCases, responseMeta, nil
}

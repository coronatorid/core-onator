package usecase

import (
	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportReject reject report provider by users
type ReportReject struct {
}

// Perform report rejection
func (r *ReportReject) Perform(ctx provider.Context, adminID, reportedCasesID int, adminProvider provider.Admin, reportProvider provider.Report) (entity.ReportedCases, *entity.ApplicationError) {
	if _, err := adminProvider.Authenticate(ctx, adminID, []constant.UserRole{constant.UserRoleSuperAdmin}); err != nil {
		return entity.ReportedCases{}, err
	}

	if err := reportProvider.Reject(ctx, reportedCasesID); err != nil {
		return entity.ReportedCases{}, err
	}

	reportedCase, err := reportProvider.Find(ctx, reportedCasesID)
	if err != nil {
		return entity.ReportedCases{}, err
	}

	return reportedCase, nil
}

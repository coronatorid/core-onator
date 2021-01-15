package usecase

import (
	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportDelete ...
type ReportDelete struct{}

// Perform delete report
func (r *ReportDelete) Perform(ctx provider.Context, adminID, reportedCasesID int, adminProvider provider.Admin, reportProvider provider.Report) *entity.ApplicationError {
	if _, err := adminProvider.Authenticate(ctx, adminID, []constant.UserRole{constant.UserRoleSuperAdmin}); err != nil {
		return err
	}

	reportedCases, err := reportProvider.Find(ctx, reportedCasesID)
	if err != nil {
		return err
	}

	return reportProvider.DeleteReportedCases(ctx, reportedCases.UserID)
}

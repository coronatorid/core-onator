package usecase

import (
	"errors"
	"net/http"

	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// ReportConfirm reject report provider by users
type ReportConfirm struct {
}

// Perform report rejection
func (r *ReportConfirm) Perform(ctx provider.Context, adminID, reportedCasesID int, adminProvider provider.Admin, reportProvider provider.Report) (entity.ReportedCases, *entity.ApplicationError) {
	if _, err := adminProvider.Authenticate(ctx, adminID, []constant.UserRole{constant.UserRoleSuperAdmin}); err != nil {
		return entity.ReportedCases{}, err
	}

	reportedCase, err := reportProvider.Find(ctx, reportedCasesID)
	if err != nil {
		return entity.ReportedCases{}, err
	}

	if reportedCase.Status == constant.ReportedCasesConfirmed.Int() {
		return entity.ReportedCases{}, &entity.ApplicationError{
			Err:        []error{errors.New("Laporan hanya bisa dikonfirmasi satu kali")},
			HTTPStatus: http.StatusUnprocessableEntity,
		}
	}

	if err := reportProvider.Confirm(ctx, reportedCasesID); err != nil {
		return entity.ReportedCases{}, err
	}

	reportedCase, err = reportProvider.Find(ctx, reportedCasesID)
	if err != nil {
		return entity.ReportedCases{}, err
	}

	return reportedCase, nil
}

package usecase

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// DeleteReportedCases ...
type DeleteReportedCases struct{}

// Perform delete reported cases
func (d *DeleteReportedCases) Perform(ctx provider.Context, userID int, reportProvider provider.Report) *entity.ApplicationError {
	reportedCases, err := reportProvider.FindByUserID(ctx, userID)
	if err != nil {

	}

	if err := reportProvider.Delete(ctx, reportedCases.ID); err != nil {
		return err
	}

	_ = reportProvider.DeleteFile(ctx, reportedCases.ImagePath)

	return nil
}

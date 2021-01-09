package usecase

import (
	"net/url"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// DeleteReportedCases ...
type DeleteReportedCases struct{}

// Perform delete reported cases
func (d *DeleteReportedCases) Perform(ctx provider.Context, userID int, reportProvider provider.Report) *entity.ApplicationError {
	reportedCases, applicationErr := reportProvider.FindByUserID(ctx, userID)
	if applicationErr != nil {
		return applicationErr
	}

	if applicationErr := reportProvider.Delete(ctx, reportedCases.ID); applicationErr != nil {
		return applicationErr
	}

	parsedURL, err := url.Parse(reportedCases.ImagePath)
	if err == nil {
		path := parsedURL.Path
		_ = reportProvider.DeleteFile(ctx, path[1:len(path)])
	}

	return nil
}

package usecase

import (
	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Confirm reported cases
type Confirm struct{}

// Perform reject reported cases
func (r *Confirm) Perform(ctx provider.Context, ID int, reportProvider provider.Report) *entity.ApplicationError {
	return reportProvider.UpdateState(ctx, constant.ReportedCasesConfirmed, ID)
}

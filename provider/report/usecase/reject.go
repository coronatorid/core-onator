package usecase

import (
	"errors"
	"net/http"

	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Reject reported cases
type Reject struct{}

// Perform reject reported cases
func (r *Reject) Perform(ctx provider.Context, ID int, reportProvider provider.Report) *entity.ApplicationError {
	reportedCase, err := reportProvider.Find(ctx, ID)
	if err != nil {
		return err
	}

	if reportedCase.Status == constant.ReportedCasesConfirmed.Int() {
		return &entity.ApplicationError{
			Err:        []error{errors.New("invalid state")},
			HTTPStatus: http.StatusUnprocessableEntity,
		}
	}

	return reportProvider.UpdateState(ctx, constant.ReportedCasesRejected, ID)
}

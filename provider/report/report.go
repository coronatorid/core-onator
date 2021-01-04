package report

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/report/usecase"
)

// Report provide function for managing user data
type Report struct {
	db provider.DB
}

// Fabricate user provider
func Fabricate(db provider.DB) *Report {
	return &Report{db: db}
}

// CreateReportCases create new reported cases data
func (r *Report) CreateReportCases(ctx provider.Context, insertable entity.ReportInsertable, tx provider.TX) (int, *entity.ApplicationError) {
	createReportCases := usecase.CreateReportCases{}
	return createReportCases.Perform(ctx, insertable, tx)
}
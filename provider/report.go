package provider

import "github.com/coronatorid/core-onator/entity"

//go:generate mockgen -source=./report.go -destination=./mocks/report_mock.go -package mockProvider

// Report provider handle all reporting logic from coronator
type Report interface {
	CreateReportCases(ctx Context, insertable entity.ReportInsertable, tx TX) (int, *entity.ApplicationError)
}

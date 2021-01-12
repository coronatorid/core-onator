package provider

import (
	"mime/multipart"

	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./report.go -destination=./mocks/report_mock.go -package mockProvider

// Report provider handle all reporting logic from coronator
type Report interface {
	Create(ctx Context, insertable entity.ReportInsertable, tx TX) (int, *entity.ApplicationError)
	Delete(ctx Context, ID int) *entity.ApplicationError
	FindByUserID(ctx Context, userID int) (entity.ReportedCases, *entity.ApplicationError)

	UploadFile(ctx Context, userID int, fileHeader *multipart.FileHeader) (string, *entity.ApplicationError)
	DeleteFile(ctx Context, filePath string) *entity.ApplicationError

	CreateReportedCases(ctx Context, userID int, fileHeader *multipart.FileHeader) *entity.ApplicationError
	DeleteReportedCases(ctx Context, userID int) *entity.ApplicationError
}

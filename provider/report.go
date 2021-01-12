package provider

import (
	"mime/multipart"

	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./report.go -destination=./mocks/report_mock.go -package mockProvider

// Report provider handle all reporting logic from coronator
type Report interface {
	Create(ctx Context, insertable entity.ReportInsertable, tx TX) (int, *entity.ApplicationError)
	Delete(ctx Context, ID int) *entity.ApplicationError
	Find(ctx Context, ID int) (entity.ReportedCases, *entity.ApplicationError)
	FindByUserID(ctx Context, userID int) (entity.ReportedCases, *entity.ApplicationError)
	Count(ctx Context, status constant.ReportedCasesStatus) (int, *entity.ApplicationError)
	List(ctx Context, status constant.ReportedCasesStatus, requestMeta entity.RequestMeta) ([]entity.ReportedCases, *entity.ApplicationError)

	UploadFile(ctx Context, userID int, fileHeader *multipart.FileHeader) (string, *entity.ApplicationError)
	DeleteFile(ctx Context, filePath string) *entity.ApplicationError

	CreateReportedCases(ctx Context, userID int, fileHeader *multipart.FileHeader) *entity.ApplicationError
	DeleteReportedCases(ctx Context, userID int) *entity.ApplicationError
}

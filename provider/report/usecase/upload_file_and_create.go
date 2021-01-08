package usecase

import (
	"mime/multipart"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// UploadFileAndCreate data for reported cases
type UploadFileAndCreate struct{}

// Perform logic to new reported cases
func (c *UploadFileAndCreate) Perform(ctx provider.Context, userID int, fileHeader *multipart.FileHeader, db provider.DB, report provider.Report) *entity.ApplicationError {
	var applicationError *entity.ApplicationError

	db.Transaction(ctx, "report/upload_file_and_create_report", func(tx provider.TX) error {
		path, err := report.UploadFile(ctx, userID, fileHeader)
		if err != nil {
			applicationError = err
			return err
		}

		_, err = report.CreateReportCases(ctx, entity.ReportInsertable{ImagePath: path, UserID: userID}, tx)
		if err != nil {
			applicationError = err
			return err
		}

		return nil
	})

	return applicationError
}

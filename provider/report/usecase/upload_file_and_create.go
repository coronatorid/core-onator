package usecase

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// UploadFileAndCreate data for reported cases
type UploadFileAndCreate struct{}

// Perform logic to new reported cases
func (c *UploadFileAndCreate) Perform(ctx provider.Context, userID int, fileHeader *multipart.FileHeader, db provider.DB, report provider.Report) *entity.ApplicationError {
	var applicationError *entity.ApplicationError

	db.Transaction(ctx, "report/upload_file_and_create_report", func(tx provider.TX) error {
		_, err := report.FindByUserID(ctx, userID)
		if err == nil {
			applicationError = &entity.ApplicationError{
				Err:        []error{errors.New("report already submitted")},
				HTTPStatus: http.StatusUnprocessableEntity,
			}
			return applicationError
		}

		path, err := report.UploadFile(ctx, userID, fileHeader)
		if err != nil {
			applicationError = err
			return err
		}

		_, err = report.Create(ctx, entity.ReportInsertable{ImagePath: path, UserID: userID}, tx)
		if err != nil {
			applicationError = err
			return err
		}

		return nil
	})

	return applicationError
}

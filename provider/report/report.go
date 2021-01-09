package report

import (
	"mime/multipart"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/report/api"
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

// FabricateAPI fabricating report related API
func (r *Report) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewReportMe(r))
	engine.InjectAPI(api.NewReportMeCreate(r))
	engine.InjectAPI(api.NewReportMeDelete(r))
}

// CreateReportCases create new reported cases data
func (r *Report) CreateReportCases(ctx provider.Context, insertable entity.ReportInsertable, tx provider.TX) (int, *entity.ApplicationError) {
	createReportCases := usecase.CreateReportCases{}
	return createReportCases.Perform(ctx, insertable, tx)
}

// UploadFile into coronator storage
func (r *Report) UploadFile(ctx provider.Context, userID int, fileHeader *multipart.FileHeader) (string, *entity.ApplicationError) {
	uploadFile := usecase.UploadFile{}
	return uploadFile.Perform(ctx, userID, fileHeader)
}

// CreateReportAndUploadFile ...
func (r *Report) CreateReportAndUploadFile(ctx provider.Context, userID int, fileHeader *multipart.FileHeader) *entity.ApplicationError {
	uploadFileAndCreate := usecase.UploadFileAndCreate{}
	return uploadFileAndCreate.Perform(ctx, userID, fileHeader, r.db, r)
}

// FindByUserID ...
func (r *Report) FindByUserID(ctx provider.Context, userID int) (entity.ReportedCases, *entity.ApplicationError) {
	findByUserID := usecase.FindByUserID{}
	return findByUserID.Perform(ctx, userID, r.db)
}

// Delete ...
func (r *Report) Delete(ctx provider.Context, ID int) *entity.ApplicationError {
	delete := usecase.Delete{}
	return delete.Perform(ctx, ID, r.db)
}

// DeleteFile ...
func (r *Report) DeleteFile(ctx provider.Context, filePath string) *entity.ApplicationError {
	deleteFile := usecase.DeleteFile{}
	return deleteFile.Perform(ctx, filePath)
}

// DeleteReportedCases ...
func (r *Report) DeleteReportedCases(ctx provider.Context, userID int) *entity.ApplicationError {
	deleteReportedCases := usecase.DeleteReportedCases{}
	return deleteReportedCases.Perform(ctx, userID, r)
}

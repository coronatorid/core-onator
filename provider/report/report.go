package report

import (
	"mime/multipart"

	"github.com/coronatorid/core-onator/constant"
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

// Create create new reported cases data
func (r *Report) Create(ctx provider.Context, insertable entity.ReportInsertable, tx provider.TX) (int, *entity.ApplicationError) {
	create := usecase.Create{}
	return create.Perform(ctx, insertable, tx)
}

// UploadFile into coronator storage
func (r *Report) UploadFile(ctx provider.Context, userID int, fileHeader *multipart.FileHeader) (string, *entity.ApplicationError) {
	uploadFile := usecase.UploadFile{}
	return uploadFile.Perform(ctx, userID, fileHeader)
}

// CreateReportedCases ...
func (r *Report) CreateReportedCases(ctx provider.Context, userID int, fileHeader *multipart.FileHeader) *entity.ApplicationError {
	uploadFileAndCreate := usecase.CreateReportedCases{}
	return uploadFileAndCreate.Perform(ctx, userID, fileHeader, r.db, r)
}

// FindByUserID ...
func (r *Report) FindByUserID(ctx provider.Context, userID int) (entity.ReportedCases, *entity.ApplicationError) {
	findByUserID := usecase.FindByUserID{}
	return findByUserID.Perform(ctx, userID, r.db)
}

// Find ...
func (r *Report) Find(ctx provider.Context, ID int) (entity.ReportedCases, *entity.ApplicationError) {
	find := usecase.Find{}
	return find.Perform(ctx, ID, r.db)
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

// Count ...
func (r *Report) Count(ctx provider.Context, status constant.ReportedCasesStatus) (int, *entity.ApplicationError) {
	count := usecase.Count{}
	return count.Perform(ctx, status, r.db)
}

// List ...
func (r *Report) List(ctx provider.Context, status constant.ReportedCasesStatus, requestMeta entity.RequestMeta) ([]entity.ReportedCases, *entity.ApplicationError) {
	list := usecase.List{}
	return list.Perform(ctx, status, requestMeta, r.db)
}

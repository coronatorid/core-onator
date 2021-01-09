// Code generated by MockGen. DO NOT EDIT.
// Source: ./report.go

// Package mockProvider is a generated GoMock package.
package mockProvider

import (
	entity "github.com/coronatorid/core-onator/entity"
	provider "github.com/coronatorid/core-onator/provider"
	gomock "github.com/golang/mock/gomock"
	multipart "mime/multipart"
	reflect "reflect"
)

// MockReport is a mock of Report interface
type MockReport struct {
	ctrl     *gomock.Controller
	recorder *MockReportMockRecorder
}

// MockReportMockRecorder is the mock recorder for MockReport
type MockReportMockRecorder struct {
	mock *MockReport
}

// NewMockReport creates a new mock instance
func NewMockReport(ctrl *gomock.Controller) *MockReport {
	mock := &MockReport{ctrl: ctrl}
	mock.recorder = &MockReportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReport) EXPECT() *MockReportMockRecorder {
	return m.recorder
}

// CreateReportCases mocks base method
func (m *MockReport) CreateReportCases(ctx provider.Context, insertable entity.ReportInsertable, tx provider.TX) (int, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReportCases", ctx, insertable, tx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// CreateReportCases indicates an expected call of CreateReportCases
func (mr *MockReportMockRecorder) CreateReportCases(ctx, insertable, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReportCases", reflect.TypeOf((*MockReport)(nil).CreateReportCases), ctx, insertable, tx)
}

// UploadFile mocks base method
func (m *MockReport) UploadFile(ctx provider.Context, userID int, fileHeader *multipart.FileHeader) (string, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", ctx, userID, fileHeader)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile
func (mr *MockReportMockRecorder) UploadFile(ctx, userID, fileHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockReport)(nil).UploadFile), ctx, userID, fileHeader)
}

// CreateReportAndUploadFile mocks base method
func (m *MockReport) CreateReportAndUploadFile(ctx provider.Context, userID int, fileHeader *multipart.FileHeader) *entity.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReportAndUploadFile", ctx, userID, fileHeader)
	ret0, _ := ret[0].(*entity.ApplicationError)
	return ret0
}

// CreateReportAndUploadFile indicates an expected call of CreateReportAndUploadFile
func (mr *MockReportMockRecorder) CreateReportAndUploadFile(ctx, userID, fileHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReportAndUploadFile", reflect.TypeOf((*MockReport)(nil).CreateReportAndUploadFile), ctx, userID, fileHeader)
}

// FindByUserID mocks base method
func (m *MockReport) FindByUserID(ctx provider.Context, userID int) (entity.ReportedCases, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", ctx, userID)
	ret0, _ := ret[0].(entity.ReportedCases)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID
func (mr *MockReportMockRecorder) FindByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockReport)(nil).FindByUserID), ctx, userID)
}

// Delete mocks base method
func (m *MockReport) Delete(ctx provider.Context, ID int) *entity.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, ID)
	ret0, _ := ret[0].(*entity.ApplicationError)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockReportMockRecorder) Delete(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReport)(nil).Delete), ctx, ID)
}

// DeleteFile mocks base method
func (m *MockReport) DeleteFile(ctx provider.Context, filePath string) *entity.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", ctx, filePath)
	ret0, _ := ret[0].(*entity.ApplicationError)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile
func (mr *MockReportMockRecorder) DeleteFile(ctx, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockReport)(nil).DeleteFile), ctx, filePath)
}

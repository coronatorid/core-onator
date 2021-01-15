// Code generated by MockGen. DO NOT EDIT.
// Source: ./report.go

// Package mockProvider is a generated GoMock package.
package mockProvider

import (
	constant "github.com/coronatorid/core-onator/constant"
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

// Create mocks base method
func (m *MockReport) Create(ctx provider.Context, insertable entity.ReportInsertable, tx provider.TX) (int, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, insertable, tx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockReportMockRecorder) Create(ctx, insertable, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockReport)(nil).Create), ctx, insertable, tx)
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

// Find mocks base method
func (m *MockReport) Find(ctx provider.Context, ID int) (entity.ReportedCases, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, ID)
	ret0, _ := ret[0].(entity.ReportedCases)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockReportMockRecorder) Find(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockReport)(nil).Find), ctx, ID)
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

// Count mocks base method
func (m *MockReport) Count(ctx provider.Context, status constant.ReportedCasesStatus) (int, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, status)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockReportMockRecorder) Count(ctx, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockReport)(nil).Count), ctx, status)
}

// List mocks base method
func (m *MockReport) List(ctx provider.Context, status constant.ReportedCasesStatus, requestMeta entity.RequestMeta) ([]entity.ReportedCases, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, status, requestMeta)
	ret0, _ := ret[0].([]entity.ReportedCases)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockReportMockRecorder) List(ctx, status, requestMeta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReport)(nil).List), ctx, status, requestMeta)
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

// CreateReportedCases mocks base method
func (m *MockReport) CreateReportedCases(ctx provider.Context, userID int, fileHeader *multipart.FileHeader) *entity.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReportedCases", ctx, userID, fileHeader)
	ret0, _ := ret[0].(*entity.ApplicationError)
	return ret0
}

// CreateReportedCases indicates an expected call of CreateReportedCases
func (mr *MockReportMockRecorder) CreateReportedCases(ctx, userID, fileHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReportedCases", reflect.TypeOf((*MockReport)(nil).CreateReportedCases), ctx, userID, fileHeader)
}

// DeleteReportedCases mocks base method
func (m *MockReport) DeleteReportedCases(ctx provider.Context, userID int) *entity.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReportedCases", ctx, userID)
	ret0, _ := ret[0].(*entity.ApplicationError)
	return ret0
}

// DeleteReportedCases indicates an expected call of DeleteReportedCases
func (mr *MockReportMockRecorder) DeleteReportedCases(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReportedCases", reflect.TypeOf((*MockReport)(nil).DeleteReportedCases), ctx, userID)
}

// UpdateState mocks base method
func (m *MockReport) UpdateState(ctx provider.Context, state constant.ReportedCasesStatus, ID int) *entity.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateState", ctx, state, ID)
	ret0, _ := ret[0].(*entity.ApplicationError)
	return ret0
}

// UpdateState indicates an expected call of UpdateState
func (mr *MockReportMockRecorder) UpdateState(ctx, state, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateState", reflect.TypeOf((*MockReport)(nil).UpdateState), ctx, state, ID)
}

// Reject mocks base method
func (m *MockReport) Reject(ctx provider.Context, ID int) *entity.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reject", ctx, ID)
	ret0, _ := ret[0].(*entity.ApplicationError)
	return ret0
}

// Reject indicates an expected call of Reject
func (mr *MockReportMockRecorder) Reject(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockReport)(nil).Reject), ctx, ID)
}

// Confirm mocks base method
func (m *MockReport) Confirm(ctx provider.Context, ID int) *entity.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Confirm", ctx, ID)
	ret0, _ := ret[0].(*entity.ApplicationError)
	return ret0
}

// Confirm indicates an expected call of Confirm
func (mr *MockReportMockRecorder) Confirm(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Confirm", reflect.TypeOf((*MockReport)(nil).Confirm), ctx, ID)
}

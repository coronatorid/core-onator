// Code generated by MockGen. DO NOT EDIT.
// Source: ./admin.go

// Package mockProvider is a generated GoMock package.
package mockProvider

import (
	constant "github.com/coronatorid/core-onator/constant"
	entity "github.com/coronatorid/core-onator/entity"
	provider "github.com/coronatorid/core-onator/provider"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAdmin is a mock of Admin interface
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockAdmin) Login(ctx provider.Context, request entity.Login) (entity.LoginResponse, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, request)
	ret0, _ := ret[0].(entity.LoginResponse)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAdminMockRecorder) Login(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAdmin)(nil).Login), ctx, request)
}

// RequestOTP mocks base method
func (m *MockAdmin) RequestOTP(ctx provider.Context, request entity.RequestOTP) (*entity.RequestOTPResponse, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestOTP", ctx, request)
	ret0, _ := ret[0].(*entity.RequestOTPResponse)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// RequestOTP indicates an expected call of RequestOTP
func (mr *MockAdminMockRecorder) RequestOTP(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestOTP", reflect.TypeOf((*MockAdmin)(nil).RequestOTP), ctx, request)
}

// Authenticate mocks base method
func (m *MockAdmin) Authenticate(ctx provider.Context, adminID int, allowedRole []constant.UserRole) (entity.User, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, adminID, allowedRole)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockAdminMockRecorder) Authenticate(ctx, adminID, allowedRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAdmin)(nil).Authenticate), ctx, adminID, allowedRole)
}

// ReportList mocks base method
func (m *MockAdmin) ReportList(ctx provider.Context, adminID int, status constant.ReportedCasesStatus, requestMeta entity.RequestMeta) ([]entity.ReportedCases, entity.ResponseMeta, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportList", ctx, adminID, status, requestMeta)
	ret0, _ := ret[0].([]entity.ReportedCases)
	ret1, _ := ret[1].(entity.ResponseMeta)
	ret2, _ := ret[2].(*entity.ApplicationError)
	return ret0, ret1, ret2
}

// ReportList indicates an expected call of ReportList
func (mr *MockAdminMockRecorder) ReportList(ctx, adminID, status, requestMeta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportList", reflect.TypeOf((*MockAdmin)(nil).ReportList), ctx, adminID, status, requestMeta)
}

// ReportDelete mocks base method
func (m *MockAdmin) ReportDelete(ctx provider.Context, adminID, reportedCasesID int) *entity.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportDelete", ctx, adminID, reportedCasesID)
	ret0, _ := ret[0].(*entity.ApplicationError)
	return ret0
}

// ReportDelete indicates an expected call of ReportDelete
func (mr *MockAdminMockRecorder) ReportDelete(ctx, adminID, reportedCasesID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportDelete", reflect.TypeOf((*MockAdmin)(nil).ReportDelete), ctx, adminID, reportedCasesID)
}

// ReportReject mocks base method
func (m *MockAdmin) ReportReject(ctx provider.Context, adminID, reportedCasesID int) (entity.ReportedCases, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportReject", ctx, adminID, reportedCasesID)
	ret0, _ := ret[0].(entity.ReportedCases)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// ReportReject indicates an expected call of ReportReject
func (mr *MockAdminMockRecorder) ReportReject(ctx, adminID, reportedCasesID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportReject", reflect.TypeOf((*MockAdmin)(nil).ReportReject), ctx, adminID, reportedCasesID)
}

// ReportConfirm mocks base method
func (m *MockAdmin) ReportConfirm(ctx provider.Context, adminID, reportedCasesID int) (entity.ReportedCases, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportConfirm", ctx, adminID, reportedCasesID)
	ret0, _ := ret[0].(entity.ReportedCases)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// ReportConfirm indicates an expected call of ReportConfirm
func (mr *MockAdminMockRecorder) ReportConfirm(ctx, adminID, reportedCasesID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportConfirm", reflect.TypeOf((*MockAdmin)(nil).ReportConfirm), ctx, adminID, reportedCasesID)
}

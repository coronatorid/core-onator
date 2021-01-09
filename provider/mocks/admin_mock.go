// Code generated by MockGen. DO NOT EDIT.
// Source: ./admin.go

// Package mockProvider is a generated GoMock package.
package mockProvider

import (
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

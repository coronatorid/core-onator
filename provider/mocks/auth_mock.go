// Code generated by MockGen. DO NOT EDIT.
// Source: ./auth.go

// Package mockProvider is a generated GoMock package.
package mockProvider

import (
	context "context"
	entity "github.com/coronatorid/core-onator/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuth is a mock of Auth interface
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// RequestOTP mocks base method
func (m *MockAuth) RequestOTP(ctx context.Context, request entity.RequestOTP) (*entity.RequestOTPResponse, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestOTP", ctx, request)
	ret0, _ := ret[0].(*entity.RequestOTPResponse)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// RequestOTP indicates an expected call of RequestOTP
func (mr *MockAuthMockRecorder) RequestOTP(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestOTP", reflect.TypeOf((*MockAuth)(nil).RequestOTP), ctx, request)
}

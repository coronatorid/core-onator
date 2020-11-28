// Code generated by MockGen. DO NOT EDIT.
// Source: ./user.go

// Package mockProvider is a generated GoMock package.
package mockProvider

import (
	context "context"
	entity "github.com/coronatorid/core-onator/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUser is a mock of User interface
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockUser) Find(ctx context.Context, ID int) (entity.User, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, ID)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUserMockRecorder) Find(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUser)(nil).Find), ctx, ID)
}

// FindByPhoneNumber mocks base method
func (m *MockUser) FindByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPhoneNumber", ctx, phoneNumber)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// FindByPhoneNumber indicates an expected call of FindByPhoneNumber
func (mr *MockUserMockRecorder) FindByPhoneNumber(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPhoneNumber", reflect.TypeOf((*MockUser)(nil).FindByPhoneNumber), ctx, phoneNumber)
}

// Create mocks base method
func (m *MockUser) Create(ctx context.Context, userInsertable entity.UserInsertable) (int, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userInsertable)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserMockRecorder) Create(ctx, userInsertable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUser)(nil).Create), ctx, userInsertable)
}

// CreateOrFind mocks base method
func (m *MockUser) CreateOrFind(ctx context.Context, phoneNumber string) (entity.User, *entity.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrFind", ctx, phoneNumber)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(*entity.ApplicationError)
	return ret0, ret1
}

// CreateOrFind indicates an expected call of CreateOrFind
func (mr *MockUserMockRecorder) CreateOrFind(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrFind", reflect.TypeOf((*MockUser)(nil).CreateOrFind), ctx, phoneNumber)
}
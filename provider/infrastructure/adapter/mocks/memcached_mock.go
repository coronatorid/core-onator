// Code generated by MockGen. DO NOT EDIT.
// Source: ./memcached.go

// Package mockAdapter is a generated GoMock package.
package mockAdapter

import (
	memcache "github.com/bradfitz/gomemcache/memcache"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMemcachedClient is a mock of MemcachedClient interface
type MockMemcachedClient struct {
	ctrl     *gomock.Controller
	recorder *MockMemcachedClientMockRecorder
}

// MockMemcachedClientMockRecorder is the mock recorder for MockMemcachedClient
type MockMemcachedClientMockRecorder struct {
	mock *MockMemcachedClient
}

// NewMockMemcachedClient creates a new mock instance
func NewMockMemcachedClient(ctrl *gomock.Controller) *MockMemcachedClient {
	mock := &MockMemcachedClient{ctrl: ctrl}
	mock.recorder = &MockMemcachedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMemcachedClient) EXPECT() *MockMemcachedClientMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockMemcachedClient) Set(item *memcache.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockMemcachedClientMockRecorder) Set(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockMemcachedClient)(nil).Set), item)
}

// Get mocks base method
func (m *MockMemcachedClient) Get(key string) (*memcache.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*memcache.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMemcachedClientMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMemcachedClient)(nil).Get), key)
}

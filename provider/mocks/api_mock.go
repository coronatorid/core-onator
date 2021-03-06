// Code generated by MockGen. DO NOT EDIT.
// Source: ./api.go

// Package mockProvider is a generated GoMock package.
package mockProvider

import (
	context "context"
	provider "github.com/coronatorid/core-onator/provider"
	gomock "github.com/golang/mock/gomock"
	multipart "mime/multipart"
	http "net/http"
	url "net/url"
	reflect "reflect"
)

// MockAPIContext is a mock of APIContext interface
type MockAPIContext struct {
	ctrl     *gomock.Controller
	recorder *MockAPIContextMockRecorder
}

// MockAPIContextMockRecorder is the mock recorder for MockAPIContext
type MockAPIContextMockRecorder struct {
	mock *MockAPIContext
}

// NewMockAPIContext creates a new mock instance
func NewMockAPIContext(ctrl *gomock.Controller) *MockAPIContext {
	mock := &MockAPIContext{ctrl: ctrl}
	mock.recorder = &MockAPIContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIContext) EXPECT() *MockAPIContextMockRecorder {
	return m.recorder
}

// Request mocks base method
func (m *MockAPIContext) Request() *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// Request indicates an expected call of Request
func (mr *MockAPIContextMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockAPIContext)(nil).Request))
}

// RealIP mocks base method
func (m *MockAPIContext) RealIP() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RealIP")
	ret0, _ := ret[0].(string)
	return ret0
}

// RealIP indicates an expected call of RealIP
func (mr *MockAPIContextMockRecorder) RealIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RealIP", reflect.TypeOf((*MockAPIContext)(nil).RealIP))
}

// Path mocks base method
func (m *MockAPIContext) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path
func (mr *MockAPIContextMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockAPIContext)(nil).Path))
}

// Param mocks base method
func (m *MockAPIContext) Param(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param
func (mr *MockAPIContextMockRecorder) Param(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockAPIContext)(nil).Param), name)
}

// ParamNames mocks base method
func (m *MockAPIContext) ParamNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParamNames indicates an expected call of ParamNames
func (mr *MockAPIContextMockRecorder) ParamNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamNames", reflect.TypeOf((*MockAPIContext)(nil).ParamNames))
}

// ParamValues mocks base method
func (m *MockAPIContext) ParamValues() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamValues")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParamValues indicates an expected call of ParamValues
func (mr *MockAPIContextMockRecorder) ParamValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamValues", reflect.TypeOf((*MockAPIContext)(nil).ParamValues))
}

// QueryParam mocks base method
func (m *MockAPIContext) QueryParam(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryParam", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// QueryParam indicates an expected call of QueryParam
func (mr *MockAPIContextMockRecorder) QueryParam(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryParam", reflect.TypeOf((*MockAPIContext)(nil).QueryParam), name)
}

// QueryParams mocks base method
func (m *MockAPIContext) QueryParams() url.Values {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryParams")
	ret0, _ := ret[0].(url.Values)
	return ret0
}

// QueryParams indicates an expected call of QueryParams
func (mr *MockAPIContextMockRecorder) QueryParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryParams", reflect.TypeOf((*MockAPIContext)(nil).QueryParams))
}

// QueryString mocks base method
func (m *MockAPIContext) QueryString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryString")
	ret0, _ := ret[0].(string)
	return ret0
}

// QueryString indicates an expected call of QueryString
func (mr *MockAPIContextMockRecorder) QueryString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryString", reflect.TypeOf((*MockAPIContext)(nil).QueryString))
}

// FormFile mocks base method
func (m *MockAPIContext) FormFile(name string) (*multipart.FileHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormFile", name)
	ret0, _ := ret[0].(*multipart.FileHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormFile indicates an expected call of FormFile
func (mr *MockAPIContextMockRecorder) FormFile(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormFile", reflect.TypeOf((*MockAPIContext)(nil).FormFile), name)
}

// File mocks base method
func (m *MockAPIContext) File(file string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "File", file)
	ret0, _ := ret[0].(error)
	return ret0
}

// File indicates an expected call of File
func (mr *MockAPIContextMockRecorder) File(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "File", reflect.TypeOf((*MockAPIContext)(nil).File), file)
}

// Cookie mocks base method
func (m *MockAPIContext) Cookie(name string) (*http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cookie", name)
	ret0, _ := ret[0].(*http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cookie indicates an expected call of Cookie
func (mr *MockAPIContextMockRecorder) Cookie(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookie", reflect.TypeOf((*MockAPIContext)(nil).Cookie), name)
}

// SetCookie mocks base method
func (m *MockAPIContext) SetCookie(cookie *http.Cookie) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCookie", cookie)
}

// SetCookie indicates an expected call of SetCookie
func (mr *MockAPIContextMockRecorder) SetCookie(cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCookie", reflect.TypeOf((*MockAPIContext)(nil).SetCookie), cookie)
}

// Cookies mocks base method
func (m *MockAPIContext) Cookies() []*http.Cookie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cookies")
	ret0, _ := ret[0].([]*http.Cookie)
	return ret0
}

// Cookies indicates an expected call of Cookies
func (mr *MockAPIContextMockRecorder) Cookies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookies", reflect.TypeOf((*MockAPIContext)(nil).Cookies))
}

// JSON mocks base method
func (m *MockAPIContext) JSON(code int, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSON", code, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSON indicates an expected call of JSON
func (mr *MockAPIContextMockRecorder) JSON(code, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockAPIContext)(nil).JSON), code, i)
}

// NoContent mocks base method
func (m *MockAPIContext) NoContent(code int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoContent", code)
	ret0, _ := ret[0].(error)
	return ret0
}

// NoContent indicates an expected call of NoContent
func (mr *MockAPIContextMockRecorder) NoContent(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoContent", reflect.TypeOf((*MockAPIContext)(nil).NoContent), code)
}

// Ctx mocks base method
func (m *MockAPIContext) Ctx() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ctx")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Ctx indicates an expected call of Ctx
func (mr *MockAPIContextMockRecorder) Ctx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ctx", reflect.TypeOf((*MockAPIContext)(nil).Ctx))
}

// Get mocks base method
func (m *MockAPIContext) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockAPIContextMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAPIContext)(nil).Get), key)
}

// Set mocks base method
func (m *MockAPIContext) Set(key string, val interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, val)
}

// Set indicates an expected call of Set
func (mr *MockAPIContextMockRecorder) Set(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockAPIContext)(nil).Set), key, val)
}

// MockAPIHandler is a mock of APIHandler interface
type MockAPIHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAPIHandlerMockRecorder
}

// MockAPIHandlerMockRecorder is the mock recorder for MockAPIHandler
type MockAPIHandlerMockRecorder struct {
	mock *MockAPIHandler
}

// NewMockAPIHandler creates a new mock instance
func NewMockAPIHandler(ctrl *gomock.Controller) *MockAPIHandler {
	mock := &MockAPIHandler{ctrl: ctrl}
	mock.recorder = &MockAPIHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIHandler) EXPECT() *MockAPIHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockAPIHandler) Handle(context provider.APIContext) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", context)
}

// Handle indicates an expected call of Handle
func (mr *MockAPIHandlerMockRecorder) Handle(context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockAPIHandler)(nil).Handle), context)
}

// Method mocks base method
func (m *MockAPIHandler) Method() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(string)
	return ret0
}

// Method indicates an expected call of Method
func (mr *MockAPIHandlerMockRecorder) Method() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockAPIHandler)(nil).Method))
}

// Path mocks base method
func (m *MockAPIHandler) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path
func (mr *MockAPIHandlerMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockAPIHandler)(nil).Path))
}

// MockAPIEngine is a mock of APIEngine interface
type MockAPIEngine struct {
	ctrl     *gomock.Controller
	recorder *MockAPIEngineMockRecorder
}

// MockAPIEngineMockRecorder is the mock recorder for MockAPIEngine
type MockAPIEngineMockRecorder struct {
	mock *MockAPIEngine
}

// NewMockAPIEngine creates a new mock instance
func NewMockAPIEngine(ctrl *gomock.Controller) *MockAPIEngine {
	mock := &MockAPIEngine{ctrl: ctrl}
	mock.recorder = &MockAPIEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIEngine) EXPECT() *MockAPIEngineMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockAPIEngine) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockAPIEngineMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockAPIEngine)(nil).Run))
}

// InjectAPI mocks base method
func (m *MockAPIEngine) InjectAPI(handler provider.APIHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InjectAPI", handler)
}

// InjectAPI indicates an expected call of InjectAPI
func (mr *MockAPIEngineMockRecorder) InjectAPI(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectAPI", reflect.TypeOf((*MockAPIEngine)(nil).InjectAPI), handler)
}

// Shutdown mocks base method
func (m *MockAPIEngine) Shutdown(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockAPIEngineMockRecorder) Shutdown(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockAPIEngine)(nil).Shutdown), ctx)
}

// MockContext is a mock of Context interface
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Ctx mocks base method
func (m *MockContext) Ctx() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ctx")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Ctx indicates an expected call of Ctx
func (mr *MockContextMockRecorder) Ctx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ctx", reflect.TypeOf((*MockContext)(nil).Ctx))
}

// Get mocks base method
func (m *MockContext) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockContextMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContext)(nil).Get), key)
}

// Set mocks base method
func (m *MockContext) Set(key string, val interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, val)
}

// Set indicates an expected call of Set
func (mr *MockContextMockRecorder) Set(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockContext)(nil).Set), key, val)
}

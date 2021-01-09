package provider

import (
	"context"
	"mime/multipart"
	"net/http"
	"net/url"
)

//go:generate mockgen -source=./api.go -destination=./mocks/api_mock.go -package mockProvider

// APIContext used by API handler to modify it's request
type APIContext interface {
	// Request returns `*http.Request`.
	Request() *http.Request

	// RealIP returns the client's network address based on `X-Forwarded-For`
	// or `X-Real-IP` request header.
	// The behavior can be configured using `Echo#IPExtractor`.
	RealIP() string

	// Path returns the registered path for the handler.
	Path() string

	// Param returns path parameter by name.
	Param(name string) string

	// ParamNames returns path parameter names.
	ParamNames() []string

	// ParamValues returns path parameter values.
	ParamValues() []string

	// QueryParam returns the query param for the provided name.
	QueryParam(name string) string

	// QueryParams returns the query parameters as `url.Values`.
	QueryParams() url.Values

	// QueryString returns the URL query string.
	QueryString() string

	// FormFile returns the multipart form file for the provided name.
	FormFile(name string) (*multipart.FileHeader, error)

	// File sends a response with the content of the file.
	File(file string) error

	// Cookie returns the named cookie provided in the request.
	Cookie(name string) (*http.Cookie, error)

	// SetCookie adds a `Set-Cookie` header in HTTP response.
	SetCookie(cookie *http.Cookie)

	// Cookies returns the HTTP cookies sent with the request.
	Cookies() []*http.Cookie

	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error

	// NoContent sends a response with no body and a status code.
	NoContent(code int) error

	// Logger returns the `Logger` instance.
	// Logger() Logger

	Context
}

// APIHandler handling api request from client
type APIHandler interface {
	Handle(context APIContext)
	Method() string
	Path() string
}

// APIEngine ...
type APIEngine interface {
	Run() error
	InjectAPI(handler APIHandler)
	Shutdown(ctx context.Context) error
}

// Context used for application flow
type Context interface {
	Ctx() context.Context

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})
}

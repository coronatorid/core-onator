package testhelper

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func PerformRequest(r http.Handler, method, path string, body io.Reader, reqModifiers ...func(req *http.Request)) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	for _, f := range reqModifiers {
		f(req)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

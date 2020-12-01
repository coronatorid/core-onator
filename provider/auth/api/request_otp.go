package api

import (
	"encoding/json"
	"net/http"

	"github.com/coronatorid/core-onator/entity"

	"github.com/coronatorid/core-onator/provider"
)

// RequestOTP api handler
type RequestOTP struct {
	authProvider provider.Auth
}

// NewRequestOTP create new request otp handler object
func NewRequestOTP(authProvider provider.Auth) *RequestOTP {
	return &RequestOTP{authProvider: authProvider}
}

// Path return api path
func (r *RequestOTP) Path() string {
	return "/authorization/otp-requests"
}

// Method return api method
func (r *RequestOTP) Method() string {
	return "POST"
}

// Handle request otp
func (r *RequestOTP) Handle(context provider.APIContext) {
	var request entity.RequestOTP
	if err := json.NewDecoder(context.Request().Body).Decode(&request); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	response, err := r.authProvider.RequestOTP(context.Request().Context(), request)
	if err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{"data": response})
}

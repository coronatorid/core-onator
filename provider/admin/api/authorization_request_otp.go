package api

import (
	"encoding/json"
	"net/http"

	"github.com/coronatorid/core-onator/entity"

	"github.com/coronatorid/core-onator/provider"
)

// AuthorizationRequestOTP api handler
type AuthorizationRequestOTP struct {
	adminProvider provider.Admin
}

// NewAuthorizationRequestOTP create new request otp handler object
func NewAuthorizationRequestOTP(adminProvider provider.Admin) *AuthorizationRequestOTP {
	return &AuthorizationRequestOTP{adminProvider: adminProvider}
}

// Path return api path
func (r *AuthorizationRequestOTP) Path() string {
	return "/administrations/authorization/otp-requests"
}

// Method return api method
func (r *AuthorizationRequestOTP) Method() string {
	return "POST"
}

// Handle request otp
func (r *AuthorizationRequestOTP) Handle(context provider.APIContext) {
	var request entity.RequestOTP
	if err := json.NewDecoder(context.Request().Body).Decode(&request); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	response, err := r.adminProvider.RequestOTP(context, request)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{"data": response})
}

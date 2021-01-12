package api

import (
	"encoding/json"
	"net/http"

	"github.com/coronatorid/core-onator/entity"

	"github.com/coronatorid/core-onator/provider"
)

// AuthorizationLogin api handler
type AuthorizationLogin struct {
	adminProvider provider.Admin
}

// NewAuthorizationLogin create new request otp handler object
func NewAuthorizationLogin(adminProvider provider.Admin) *AuthorizationLogin {
	return &AuthorizationLogin{adminProvider: adminProvider}
}

// Path return api path
func (r *AuthorizationLogin) Path() string {
	return "/administrations/authorization/login"
}

// Method return api method
func (r *AuthorizationLogin) Method() string {
	return "POST"
}

// Handle request otp
func (r *AuthorizationLogin) Handle(context provider.APIContext) {
	var request entity.Login
	if err := json.NewDecoder(context.Request().Body).Decode(&request); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	response, err := r.adminProvider.Login(context, request)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{"data": response})
}

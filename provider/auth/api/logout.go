package api

import (
	"encoding/json"
	"net/http"

	"github.com/coronatorid/core-onator/entity"

	"github.com/coronatorid/core-onator/provider"
)

// Logout api handler
type Logout struct {
	authProvider provider.Auth
}

// NewLogout create new request otp handler object
func NewLogout(authProvider provider.Auth) *Logout {
	return &Logout{authProvider: authProvider}
}

// Path return api path
func (r *Logout) Path() string {
	return "/authorization/logout"
}

// Method return api method
func (r *Logout) Method() string {
	return "POST"
}

// Handle request otp
func (r *Logout) Handle(context provider.APIContext) {
	var request entity.RevokeTokenRequest
	if err := json.NewDecoder(context.Request().Body).Decode(&request); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	err := r.authProvider.Logout(context, request)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}

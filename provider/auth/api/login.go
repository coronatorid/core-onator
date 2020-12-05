package api

import (
	"encoding/json"
	"net/http"

	"github.com/coronatorid/core-onator/entity"

	"github.com/coronatorid/core-onator/provider"
)

// Login api handler
type Login struct {
	authProvider provider.Auth
}

// NewLogin create new request otp handler object
func NewLogin(authProvider provider.Auth) *Login {
	return &Login{authProvider: authProvider}
}

// Path return api path
func (r *Login) Path() string {
	return "/authorization/login"
}

// Method return api method
func (r *Login) Method() string {
	return "POST"
}

// Handle request otp
func (r *Login) Handle(context provider.APIContext) {
	var request entity.Login
	if err := json.NewDecoder(context.Request().Body).Decode(&request); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []entity.APIError{entity.ErrorBadRequest()},
			"message": "Bad request",
		})
		return
	}

	response, err := r.authProvider.Login(context.Request().Context(), request)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{"data": response})
}

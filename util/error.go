package util

import (
	"fmt"
	"net/http"

	"github.com/coronatorid/core-onator/provider"

	"github.com/coronatorid/core-onator/entity"
)

// CreateInternalServerError ...
func CreateInternalServerError(ctx provider.Context) *entity.ApplicationError {
	return &entity.ApplicationError{
		Err:        []error{fmt.Errorf("Nampaknya terjadi kesalahan pada server coronator nih :(, sampaikan code ini pada twitter @coronatorid ya biar di cek admin: %s", GetRequestID(ctx))},
		HTTPStatus: http.StatusInternalServerError,
	}
}

// CreateServiceUnavailable ...
func CreateServiceUnavailable(ctx provider.Context) *entity.ApplicationError {
	return &entity.ApplicationError{
		Err:        []error{fmt.Errorf("Nampaknya terjadi kesalahan pada server coronator nih :(, sampaikan code ini pada twitter @coronatorid ya biar di cek admin: %s", GetRequestID(ctx))},
		HTTPStatus: http.StatusServiceUnavailable,
	}
}

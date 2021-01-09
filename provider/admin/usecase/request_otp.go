package usecase

import (
	"errors"
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// RequestOTP with admin flow
type RequestOTP struct {
}

// Perform request otp with admin flow
func (r *RequestOTP) Perform(ctx provider.Context, request entity.RequestOTP, userProvider provider.User, authProvider provider.Auth) (*entity.RequestOTPResponse, *entity.ApplicationError) {
	user, err := userProvider.FindByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if user.Role == 0 {
		return nil, &entity.ApplicationError{
			Err:        []error{errors.New("only admin can use this feature")},
			HTTPStatus: http.StatusForbidden,
		}
	}

	return authProvider.RequestOTP(ctx, request, 6)
}

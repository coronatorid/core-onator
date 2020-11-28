package usecase

import (
	"context"
	"encoding/base32"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Login handle core-onator login process
type Login struct{}

// Perform login process
func (l *Login) Perform(ctx context.Context, request entity.Login, otpRetryDuration time.Duration, userProvider provider.User) *entity.ApplicationError {
	valid, err := totp.ValidateCustom(request.OTPCode, base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%sX%s", os.Getenv("OTP_SECRET"), request.PhoneNumber))), request.OTPSentTime, totp.ValidateOpts{
		Algorithm: otp.AlgorithmSHA512,
		Digits:    4,
		Period:    uint(otpRetryDuration.Seconds()),
	})
	if err != nil {
		return l.invalidOtpError()
	} else if valid == false {
		return l.invalidOtpError()
	}

	_, errProvider := userProvider.CreateOrFind(ctx, request.PhoneNumber)
	if errProvider != nil {
		return errProvider
	}

	// TODO Altair http call to get oauth token

	return nil
}

func (l *Login) invalidOtpError() *entity.ApplicationError {
	return &entity.ApplicationError{
		Err:        []error{errors.New("otp code is invalid")},
		HTTPStatus: http.StatusUnprocessableEntity,
	}
}

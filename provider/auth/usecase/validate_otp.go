package usecase

import (
	"encoding/base32"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ValidateOTP logic
type ValidateOTP struct{}

// Perform validate otp logic
func (v *ValidateOTP) Perform(ctx provider.Context, request entity.Login, otpRetryDuration time.Duration, otpDigit int) *entity.ApplicationError {
	if time.Since(request.OTPSentTime.UTC()).Seconds() > otpRetryDuration.Seconds() {
		return v.invalidOtpError()
	}

	valid, err := totp.ValidateCustom(request.OTPCode, base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%sX%s", os.Getenv("OTP_SECRET"), request.PhoneNumber))), request.OTPSentTime.UTC(), totp.ValidateOpts{
		Algorithm: otp.AlgorithmSHA512,
		Digits:    otp.Digits(otpDigit),
		Period:    uint(otpRetryDuration.Seconds()),
	})
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("auth").Str("login")).
			Msg("error when validating otp")
		return v.invalidOtpError()
	} else if valid == false {
		return v.invalidOtpError()
	}

	return nil
}

func (v *ValidateOTP) invalidOtpError() *entity.ApplicationError {
	return &entity.ApplicationError{
		Err:        []error{errors.New("otp code is invalid")},
		HTTPStatus: http.StatusUnprocessableEntity,
	}
}

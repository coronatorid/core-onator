package usecase

import (
	"encoding/base32"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// OTPMessage template
const OTPMessage = "JANGAN BERIKAN KODE OTP INI KE SIAPAPUN, TERMASUK KE PIHAK CORONATOR SENDIRI. KODE OTP INI SANGAT RAHASIA DAN DIGUNAKAN UNTUK MASUK KEDALAM APLIKASI CORONATOR. BERIKUT ADALAH KODE OTP ANDA: %s"

// RequestOTP use for sending otp to client
// Regex indonesian phone number: ^\+62\d{10,12}
type RequestOTP struct{}

// Perform otp request process
func (r *RequestOTP) Perform(ctx provider.Context, request entity.RequestOTP, regex *regexp.Regexp, cache provider.Cache, textPublisher provider.TextPublisher, otpRetryDuration time.Duration) (*entity.RequestOTPResponse, *entity.ApplicationError) {
	if err := r.validation(request, regex); err != nil {
		return nil, err
	}

	if err := r.latestRequestCache(ctx, request, cache, otpRetryDuration); err != nil {
		return nil, err
	}

	otpResponse := r.setLatestRequestCache(ctx, request, cache)
	otpCode, err := totp.GenerateCodeCustom(base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%sX%s", os.Getenv("OTP_SECRET"), request.PhoneNumber))), otpResponse.SentTime, totp.ValidateOpts{
		Algorithm: otp.AlgorithmSHA512,
		Digits:    4,
		Period:    uint(otpRetryDuration.Seconds()),
	})
	if err != nil {
		return nil, &entity.ApplicationError{
			Err:        []error{errors.New("There are error when generating otp")},
			HTTPStatus: http.StatusInternalServerError,
		}
	}

	if err := textPublisher.Publish(ctx, request.PhoneNumber, fmt.Sprintf(OTPMessage, otpCode)); err != nil {
		return nil, &entity.ApplicationError{
			Err:        []error{errors.New("Error when sending otp to whatsapp")},
			HTTPStatus: http.StatusInternalServerError,
		}
	}

	return &otpResponse, nil
}

func (r *RequestOTP) validation(request entity.RequestOTP, regex *regexp.Regexp) *entity.ApplicationError {
	if request.PhoneNumber == "" || len(request.PhoneNumber) > 14 || regex.MatchString(request.PhoneNumber) == false {
		return &entity.ApplicationError{
			Err:        []error{errors.New("Phone number request is invalid, make sure it's not empty, the length is less than 12 and use Indonesian phone number")},
			HTTPStatus: http.StatusUnprocessableEntity,
		}
	}

	return nil
}

func (r *RequestOTP) setLatestRequestCache(ctx provider.Context, request entity.RequestOTP, cache provider.Cache) entity.RequestOTPResponse {
	otpResponse := entity.RequestOTPResponse{
		PhoneNumber: request.PhoneNumber,
		SentTime:    time.Now().UTC(),
	}
	encodedOTPResponse, _ := json.Marshal(otpResponse)
	_ = cache.Set(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber), encodedOTPResponse, 0)

	return otpResponse
}

func (r *RequestOTP) latestRequestCache(ctx provider.Context, request entity.RequestOTP, cache provider.Cache, otpRetryDuration time.Duration) *entity.ApplicationError {
	item, err := cache.Get(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber))
	if err == nil {
		var lastCacheRequest entity.RequestOTPResponse
		err := json.Unmarshal(item.Value(), &lastCacheRequest)
		if subtractedTime := time.Now().UTC().Sub(lastCacheRequest.SentTime); err == nil && subtractedTime < otpRetryDuration {
			return &entity.ApplicationError{
				Err:        []error{fmt.Errorf("Please retry in %d seconds", int(otpRetryDuration.Seconds()-subtractedTime.Seconds()))},
				HTTPStatus: http.StatusTooEarly,
			}
		}
	}

	return nil
}

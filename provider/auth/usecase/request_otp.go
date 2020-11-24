package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// RequestOTP use for sending otp to client
// Regex indonesian phone number: ^\+62\d{10,12}
type RequestOTP struct{}

// Perform otp request process
func (r *RequestOTP) Perform(ctx context.Context, request entity.RequestOTP, regex *regexp.Regexp, cache provider.Cache, otpRetryDuration time.Duration) (*entity.RequestOTPResponse, *entity.ApplicationError) {
	if err := r.validation(request, regex); err != nil {
		return nil, err
	}

	if err := r.latestRequestCache(ctx, request, cache, otpRetryDuration); err != nil {
		return nil, err
	}

	// TODO should add send whatsapp in here

	otpResponse := r.setLatestRequestCache(ctx, request, cache)
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

func (r *RequestOTP) setLatestRequestCache(ctx context.Context, request entity.RequestOTP, cache provider.Cache) entity.RequestOTPResponse {
	otpResponse := entity.RequestOTPResponse{
		PhoneNumber: request.PhoneNumber,
		SentTime:    time.Now(),
	}
	encodedOTPResponse, _ := json.Marshal(otpResponse)
	_ = cache.Set(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber), encodedOTPResponse, 0)

	return otpResponse
}

func (r *RequestOTP) latestRequestCache(ctx context.Context, request entity.RequestOTP, cache provider.Cache, otpRetryDuration time.Duration) *entity.ApplicationError {
	item, err := cache.Get(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber))
	if err == nil {
		var lastCacheRequest entity.RequestOTPResponse
		err := json.Unmarshal(item.Value(), &lastCacheRequest)
		if subtractedTime := time.Now().Sub(lastCacheRequest.SentTime); err == nil && subtractedTime < otpRetryDuration {
			return &entity.ApplicationError{
				Err:        []error{fmt.Errorf("Please retry in %d seconds", int(subtractedTime.Seconds()))},
				HTTPStatus: http.StatusTooEarly,
			}
		}
	}

	return nil
}

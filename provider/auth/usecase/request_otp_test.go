package usecase_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/auth/usecase"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/coronatorid/core-onator/util"

	"github.com/coronatorid/core-onator/entity"
	"github.com/golang/mock/gomock"
)

func TestRequestOTP(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := testhelper.NewTestContext()
	regexIndonesianPhoneNumber, _ := regexp.Compile(`^\+62\d{10,12}`)
	otpRetryDuration := 30 * time.Second

	t.Run("Perform", func(t *testing.T) {
		t.Run("When performed it will send otp", func(t *testing.T) {
			request := entity.RequestOTP{
				// Random phone number
				PhoneNumber: "+6289762562712",
			}

			cache := mockProvider.NewMockCache(mockCtrl)
			cache.EXPECT().Get(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber)).Return(nil, provider.ErrCacheMiss)
			cache.EXPECT().Set(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber), gomock.Any(), 0*time.Second).DoAndReturn(func(ctx provider.Context, key string, value []byte, expiration time.Duration) error {
				var response entity.RequestOTPResponse
				err := json.Unmarshal(value, &response)

				assert.Nil(t, err)
				assert.Equal(t, request.PhoneNumber, response.PhoneNumber)

				return nil
			})

			textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)
			textPublisher.EXPECT().Publish(ctx, request.PhoneNumber, gomock.Any()).Return(nil)

			uc := usecase.RequestOTP{}
			response, err := uc.Perform(ctx, request, regexIndonesianPhoneNumber, cache, textPublisher, otpRetryDuration)

			assert.Nil(t, err)
			assert.Equal(t, request.PhoneNumber, response.PhoneNumber)
		})

		t.Run("When there is a cache but the cache is greater than 30 seconds then it will send otp", func(t *testing.T) {
			request := entity.RequestOTP{
				// Random phone number
				PhoneNumber: "+6289762562712",
			}

			cachedResponse := entity.RequestOTPResponse{PhoneNumber: request.PhoneNumber, SentTime: time.Now().UTC().Add(-60 * time.Second)}
			encodedCachedResponse, _ := json.Marshal(cachedResponse)

			cacheItem := mockProvider.NewMockCacheItem(mockCtrl)
			cacheItem.EXPECT().Value().Return(encodedCachedResponse)

			cache := mockProvider.NewMockCache(mockCtrl)
			cache.EXPECT().Get(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber)).Return(cacheItem, nil)
			cache.EXPECT().Set(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber), gomock.Any(), 0*time.Second).DoAndReturn(func(ctx provider.Context, key string, value []byte, expiration time.Duration) error {
				var response entity.RequestOTPResponse
				err := json.Unmarshal(value, &response)

				assert.Nil(t, err)
				assert.Equal(t, request.PhoneNumber, response.PhoneNumber)

				return nil
			})

			textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)
			textPublisher.EXPECT().Publish(ctx, request.PhoneNumber, gomock.Any()).Return(nil)

			uc := usecase.RequestOTP{}
			response, err := uc.Perform(ctx, request, regexIndonesianPhoneNumber, cache, textPublisher, otpRetryDuration)

			assert.Nil(t, err)
			assert.Equal(t, request.PhoneNumber, response.PhoneNumber)
		})

		t.Run("When there is a cache but the cache is less than 30 seconds then it will not send otp", func(t *testing.T) {
			request := entity.RequestOTP{
				// Random phone number
				PhoneNumber: "+6289762562712",
			}

			cachedResponse := entity.RequestOTPResponse{PhoneNumber: request.PhoneNumber, SentTime: time.Now().UTC().Add(-20 * time.Second)}
			encodedCachedResponse, _ := json.Marshal(cachedResponse)

			cacheItem := mockProvider.NewMockCacheItem(mockCtrl)
			cacheItem.EXPECT().Value().Return(encodedCachedResponse)

			cache := mockProvider.NewMockCache(mockCtrl)
			cache.EXPECT().Get(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber)).Return(cacheItem, nil)

			textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)

			uc := usecase.RequestOTP{}
			response, err := uc.Perform(ctx, request, regexIndonesianPhoneNumber, cache, textPublisher, otpRetryDuration)

			expectedError := &entity.ApplicationError{
				HTTPStatus: http.StatusTooEarly,
			}

			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Nil(t, response)
		})

		t.Run("When performed but text publisher return error then it will not send otp", func(t *testing.T) {
			request := entity.RequestOTP{
				// Random phone number
				PhoneNumber: "+6289762562712",
			}

			cache := mockProvider.NewMockCache(mockCtrl)
			cache.EXPECT().Get(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber)).Return(nil, provider.ErrCacheMiss)
			cache.EXPECT().Set(ctx, fmt.Sprintf("last-otp-request-%s", request.PhoneNumber), gomock.Any(), 0*time.Second).DoAndReturn(func(ctx provider.Context, key string, value []byte, expiration time.Duration) error {
				var response entity.RequestOTPResponse
				err := json.Unmarshal(value, &response)

				assert.Nil(t, err)
				assert.Equal(t, request.PhoneNumber, response.PhoneNumber)

				return nil
			})

			textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)
			textPublisher.EXPECT().Publish(ctx, request.PhoneNumber, gomock.Any()).Return(errors.New("unexpected error"))

			uc := usecase.RequestOTP{}
			response, err := uc.Perform(ctx, request, regexIndonesianPhoneNumber, cache, textPublisher, otpRetryDuration)

			expectedError := util.CreateInternalServerError(ctx)

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Nil(t, response)
		})

		t.Run("When request phone number is empty string then it return 422 error", func(t *testing.T) {
			request := entity.RequestOTP{
				// Random phone number
				PhoneNumber: "",
			}

			cache := mockProvider.NewMockCache(mockCtrl)
			textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)

			uc := usecase.RequestOTP{}
			response, err := uc.Perform(ctx, request, regexIndonesianPhoneNumber, cache, textPublisher, otpRetryDuration)

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("Phone number request is invalid, make sure it's not empty, the length is less than 12 and use Indonesian phone number")},
				HTTPStatus: http.StatusUnprocessableEntity,
			}

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Nil(t, response)
		})

		t.Run("When request phone number length is greater than 14 then it return 422 error", func(t *testing.T) {
			request := entity.RequestOTP{
				// Random phone number
				PhoneNumber: "+6289762562712222",
			}

			cache := mockProvider.NewMockCache(mockCtrl)
			textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)

			uc := usecase.RequestOTP{}
			response, err := uc.Perform(ctx, request, regexIndonesianPhoneNumber, cache, textPublisher, otpRetryDuration)

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("Phone number request is invalid, make sure it's not empty, the length is less than 12 and use Indonesian phone number")},
				HTTPStatus: http.StatusUnprocessableEntity,
			}

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Nil(t, response)
		})

		t.Run("When request phone number doensn't match indonesian phone number regex it retun 422 error", func(t *testing.T) {
			request := entity.RequestOTP{
				// Random phone number
				PhoneNumber: "+128998712829",
			}

			cache := mockProvider.NewMockCache(mockCtrl)
			textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)

			uc := usecase.RequestOTP{}
			response, err := uc.Perform(ctx, request, regexIndonesianPhoneNumber, cache, textPublisher, otpRetryDuration)

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("Phone number request is invalid, make sure it's not empty, the length is less than 12 and use Indonesian phone number")},
				HTTPStatus: http.StatusUnprocessableEntity,
			}

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Nil(t, response)
		})
	})
}

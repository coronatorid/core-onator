package auth_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/auth"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestAuthFabricate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cache := mockProvider.NewMockCache(mockCtrl)
	textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)
	userProvider := mockProvider.NewMockUser(mockCtrl)
	altair := mockProvider.NewMockAltair(mockCtrl)
	t.Run("Fabricate", func(t *testing.T) {
		t.Run("When everything is okay it will not return error", func(t *testing.T) {
			_ = os.Setenv("OTP_RETRY_DURATION", "30s")
			_, err := auth.Fabricate(cache, textPublisher, userProvider, altair)

			assert.Nil(t, err)
		})
	})
}

func TestAuthFabricateFailParseDuration(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cache := mockProvider.NewMockCache(mockCtrl)
	textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)
	userProvider := mockProvider.NewMockUser(mockCtrl)
	altair := mockProvider.NewMockAltair(mockCtrl)
	t.Run("Fabricate", func(t *testing.T) {
		t.Run("When duration is invalid then it return error", func(t *testing.T) {
			_ = os.Setenv("OTP_RETRY_DURATION", "abc")
			_, err := auth.Fabricate(cache, textPublisher, userProvider, altair)

			assert.NotNil(t, err)
		})
	})

}

func TestAuth(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("InjectAPI", func(t *testing.T) {
		apiEngine := mockProvider.NewMockAPIEngine(mockCtrl)
		apiEngine.EXPECT().InjectAPI(gomock.Any()).Times(2)

		cache := mockProvider.NewMockCache(mockCtrl)
		textPublisher := mockProvider.NewMockTextPublisher(mockCtrl)
		userProvider := mockProvider.NewMockUser(mockCtrl)
		altair := mockProvider.NewMockAltair(mockCtrl)

		_ = os.Setenv("OTP_RETRY_DURATION", "30s")
		authProvider, _ := auth.Fabricate(cache, textPublisher, userProvider, altair)
		authProvider.FabricateAPI(apiEngine)
	})

	t.Run("RequestOTP", func(t *testing.T) {
		t.Run("Normal Scenario", func(t *testing.T) {
			request := entity.RequestOTP{
				// Random phone number
				PhoneNumber: "+6289762562712",
			}

			ctx := testhelper.NewTestContext()

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

			userProvider := mockProvider.NewMockUser(mockCtrl)
			altair := mockProvider.NewMockAltair(mockCtrl)

			_ = os.Setenv("OTP_RETRY_DURATION", "30s")
			authProvider, _ := auth.Fabricate(cache, textPublisher, userProvider, altair)

			response, err := authProvider.RequestOTP(ctx, request)

			assert.Nil(t, err)
			assert.Equal(t, request.PhoneNumber, response.PhoneNumber)
		})
	})
}

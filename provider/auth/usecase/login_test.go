package usecase_test

import (
	"encoding/base32"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider/auth/usecase"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/golang/mock/gomock"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := testhelper.NewTestContext()
	otpRetryDuration := time.Second * 30
	otpSecret := "rahasia"

	os.Setenv("OTP_SECRET", otpSecret)

	t.Run("Perform", func(t *testing.T) {
		t.Run("When request login success it will return login response", func(t *testing.T) {
			request := entity.Login{
				ClientSecret: "secret",
				ClientUID:    "uid",
				OTPSentTime:  time.Now().UTC(),
				PhoneNumber:  "+6287609870987",
			}

			user := entity.User{
				ID: 99,
			}

			oauthAccessToken := entity.OauthAccessToken{}
			oauthAccessToken.Data.ExpiresIn = 3000

			otpCode, _ := totp.GenerateCodeCustom(base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%sX%s", otpSecret, request.PhoneNumber))), request.OTPSentTime, totp.ValidateOpts{
				Algorithm: otp.AlgorithmSHA512,
				Digits:    4,
				Period:    uint(otpRetryDuration.Seconds()),
			})

			request.OTPCode = otpCode

			userProvider := mockProvider.NewMockUser(mockCtrl)
			userProvider.EXPECT().CreateOrFind(ctx, request.PhoneNumber).Return(user, nil)

			altair := mockProvider.NewMockAltair(mockCtrl)
			altair.EXPECT().GrantToken(ctx, entity.GrantTokenRequest{
				ResourceOwnerID: user.ID,
				ResponseType:    "token",
				Scopes:          "users",
				ClientUID:       request.ClientUID,
				ClientSecret:    request.ClientSecret,
				RedirectURI:     "http://localhost:2019",
			}).Return(oauthAccessToken, nil)

			login := &usecase.Login{}
			_, err := login.Perform(ctx, request, otpRetryDuration, userProvider, altair)
			assert.Nil(t, err)
		})

		t.Run("When send otp request time is greater than retry duration then it will return invalid otp error", func(t *testing.T) {
			request := entity.Login{
				ClientSecret: "secret",
				ClientUID:    "uid",
				OTPSentTime:  time.Now().UTC().Add(-time.Second * 50),
				PhoneNumber:  "+6287609870987",
			}

			oauthAccessToken := entity.OauthAccessToken{}
			oauthAccessToken.Data.ExpiresIn = 3000

			otpCode, _ := totp.GenerateCodeCustom(base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%sX%s", otpSecret, request.PhoneNumber))), request.OTPSentTime, totp.ValidateOpts{
				Algorithm: otp.AlgorithmSHA512,
				Digits:    4,
				Period:    uint(otpRetryDuration.Seconds()),
			})

			request.OTPCode = otpCode

			userProvider := mockProvider.NewMockUser(mockCtrl)
			altair := mockProvider.NewMockAltair(mockCtrl)

			login := &usecase.Login{}

			expectedErr := &entity.ApplicationError{
				Err:        []error{errors.New("otp code is invalid")},
				HTTPStatus: http.StatusUnprocessableEntity,
			}

			_, err := login.Perform(ctx, request, otpRetryDuration, userProvider, altair)
			assert.Equal(t, expectedErr.Error(), err.Error())
			assert.Equal(t, expectedErr.HTTPStatus, err.HTTPStatus)
		})

		t.Run("When OTP is invalid it will return error", func(t *testing.T) {
			otpSecret := "xxx"

			request := entity.Login{
				ClientSecret: "secret",
				ClientUID:    "uid",
				OTPSentTime:  time.Now().UTC(),
				PhoneNumber:  "+6287609870987",
			}

			otpCode, _ := totp.GenerateCodeCustom(base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%sX%sT%s", otpSecret, request.PhoneNumber, request.OTPSentTime.Format(time.RFC3339)))), request.OTPSentTime, totp.ValidateOpts{
				Algorithm: otp.AlgorithmSHA512,
				Digits:    4,
				Period:    uint(otpRetryDuration.Seconds()),
			})

			request.OTPCode = otpCode

			userProvider := mockProvider.NewMockUser(mockCtrl)

			altair := mockProvider.NewMockAltair(mockCtrl)

			login := &usecase.Login{}

			expectedErr := &entity.ApplicationError{
				Err:        []error{errors.New("otp code is invalid")},
				HTTPStatus: http.StatusUnprocessableEntity,
			}

			_, err := login.Perform(ctx, request, otpRetryDuration, userProvider, altair)
			assert.Equal(t, expectedErr.Error(), err.Error())
			assert.Equal(t, expectedErr.HTTPStatus, err.HTTPStatus)
		})

		t.Run("When create or find user failed it will return error", func(t *testing.T) {

		})

		t.Run("When grant token failed it will return error", func(t *testing.T) {

		})
	})
}

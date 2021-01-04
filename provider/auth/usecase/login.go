package usecase

import (
	"encoding/base32"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
)

// Login handle core-onator login process
type Login struct{}

// Perform login process
func (l *Login) Perform(ctx provider.Context, request entity.Login, otpRetryDuration time.Duration, userProvider provider.User, altair provider.Altair) (entity.LoginResponse, *entity.ApplicationError) {
	var loginResponse entity.LoginResponse

	if time.Since(request.OTPSentTime.UTC()).Seconds() > otpRetryDuration.Seconds() {
		return entity.LoginResponse{}, l.invalidOtpError()
	}

	valid, err := totp.ValidateCustom(request.OTPCode, base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%sX%s", os.Getenv("OTP_SECRET"), request.PhoneNumber))), request.OTPSentTime.UTC(), totp.ValidateOpts{
		Algorithm: otp.AlgorithmSHA512,
		Digits:    4,
		Period:    uint(otpRetryDuration.Seconds()),
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("auth").Str("login")).
			Msg("error when validating otp")
		return entity.LoginResponse{}, l.invalidOtpError()
	} else if valid == false {
		return entity.LoginResponse{}, l.invalidOtpError()
	}

	user, errProvider := userProvider.CreateOrFind(ctx, request.PhoneNumber)
	if errProvider != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("auth").Str("login")).
			Msg("error when create or find user")
		return entity.LoginResponse{}, errProvider
	}

	oauthAccessToken, entityError := altair.GrantToken(ctx, entity.GrantTokenRequest{
		ResourceOwnerID: user.ID,
		ResponseType:    "token",
		Scopes:          "users",
		ClientUID:       request.ClientUID,
		ClientSecret:    request.ClientSecret,
		RedirectURI:     "http://localhost:2019",
	})
	if entityError != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("auth").Str("login")).
			Msg("error when granting access token")
		return entity.LoginResponse{}, entityError
	}

	loginResponse.User = user
	loginResponse.Auth.ExpiresIn = oauthAccessToken.Data.ExpiresIn
	loginResponse.Auth.Scopes = oauthAccessToken.Data.Scopes
	loginResponse.Auth.Token = oauthAccessToken.Data.Token
	return loginResponse, nil
}

func (l *Login) invalidOtpError() *entity.ApplicationError {
	return &entity.ApplicationError{
		Err:        []error{errors.New("otp code is invalid")},
		HTTPStatus: http.StatusUnprocessableEntity,
	}
}

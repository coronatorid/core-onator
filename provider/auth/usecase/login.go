package usecase

import (
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
)

// Login handle core-onator login process
type Login struct{}

// Perform login process
func (l *Login) Perform(ctx provider.Context, request entity.Login, otpRetryDuration time.Duration, userProvider provider.User, altair provider.Altair, otpDigit int, authProvider provider.Auth) (entity.LoginResponse, *entity.ApplicationError) {
	var loginResponse entity.LoginResponse

	if applicationErr := authProvider.ValidateOTP(ctx, request, otpDigit); applicationErr != nil {
		return entity.LoginResponse{}, applicationErr
	}

	user, errProvider := userProvider.CreateOrFind(ctx, request.PhoneNumber)
	if errProvider != nil {
		log.Error().
			Err(errProvider).
			Stack().
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
			Err(entityError).
			Stack().
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

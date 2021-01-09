package usecase

import (
	"errors"
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Login with admin flow
type Login struct{}

// Perform login with admin flow
func (l *Login) Perform(ctx provider.Context, request entity.Login, userProvider provider.User, altair provider.Altair, authProvider provider.Auth) (entity.LoginResponse, *entity.ApplicationError) {
	var loginResponse entity.LoginResponse

	user, err := userProvider.FindByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		return loginResponse, err
	}

	if user.Role == 0 {
		return loginResponse, &entity.ApplicationError{
			Err:        []error{errors.New("only admin can use this feature")},
			HTTPStatus: http.StatusForbidden,
		}
	}

	if err := authProvider.ValidateOTP(ctx, request, 6); err != nil {
		return loginResponse, err
	}

	oauthAccessToken, err := altair.GrantToken(ctx, entity.GrantTokenRequest{
		ResourceOwnerID: user.ID,
		ResponseType:    "token",
		Scopes:          "users admin",
		ClientUID:       request.ClientUID,
		ClientSecret:    request.ClientSecret,
		RedirectURI:     "http://localhost:2019",
	})
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("auth").Str("login")).
			Msg("error when granting access token")
		return loginResponse, err
	}

	loginResponse.User = user
	loginResponse.Auth.ExpiresIn = oauthAccessToken.Data.ExpiresIn
	loginResponse.Auth.Scopes = oauthAccessToken.Data.Scopes
	loginResponse.Auth.Token = oauthAccessToken.Data.Token
	return loginResponse, nil
}

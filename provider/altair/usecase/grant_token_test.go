package usecase_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/altair/usecase"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestGrantToken(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	t.Run("Perform", func(t *testing.T) {
		t.Run("When request post to altair success then it will return entity.OauthAccessToken", func(t *testing.T) {

			request := entity.GrantTokenRequest{
				RedirectURI:     "http://localhost:8080",
				ClientUID:       "client-uid",
				ClientSecret:    "secret",
				ResponseType:    "token",
				Scopes:          "users",
				ResourceOwnerID: 99,
			}

			expectedAccessToken := entity.OauthAccessToken{
				CreatedAt:          time.Now(),
				ExpiresIn:          3000,
				ID:                 1,
				OauthApplicationID: 1,
				RedirectURI:        request.RedirectURI,
				ResourceOwnerID:    request.ResourceOwnerID,
				Scopes:             request.Scopes,
				Token:              "some secret token",
			}

			networkCfg := mockProvider.NewMockNetworkConfig(mockCtrl)

			encodedJSON, _ := json.Marshal(request)
			network := mockProvider.NewMockNetwork(mockCtrl)
			network.EXPECT().POST(ctx, networkCfg, "/_plugins/oauth/authorizations", bytes.NewBuffer(encodedJSON), gomock.Any(), nil).DoAndReturn(func(ctx context.Context, cfg provider.NetworkConfig, path string, body io.Reader, successBinder interface{}, failedBinder interface{}) error {
				binder := successBinder.(*entity.OauthAccessToken)
				binder.CreatedAt = expectedAccessToken.CreatedAt
				binder.ExpiresIn = expectedAccessToken.ExpiresIn
				binder.ID = expectedAccessToken.ID
				binder.OauthApplicationID = expectedAccessToken.OauthApplicationID
				binder.RedirectURI = expectedAccessToken.RedirectURI
				binder.ResourceOwnerID = expectedAccessToken.ResourceOwnerID
				binder.Scopes = expectedAccessToken.Scopes
				binder.Token = expectedAccessToken.Token
				return nil
			})

			grantToken := usecase.GrantToken{}
			oauthAccessToken, err := grantToken.Perform(ctx, request, networkCfg, network)

			assert.Nil(t, err)
			assert.Equal(t, expectedAccessToken.CreatedAt, oauthAccessToken.CreatedAt)
			assert.Equal(t, expectedAccessToken.ExpiresIn, oauthAccessToken.ExpiresIn)
			assert.Equal(t, expectedAccessToken.ID, oauthAccessToken.ID)
			assert.Equal(t, expectedAccessToken.OauthApplicationID, oauthAccessToken.OauthApplicationID)
			assert.Equal(t, expectedAccessToken.RedirectURI, oauthAccessToken.RedirectURI)
			assert.Equal(t, expectedAccessToken.ResourceOwnerID, oauthAccessToken.ResourceOwnerID)
			assert.Equal(t, expectedAccessToken.Scopes, oauthAccessToken.Scopes)
			assert.Equal(t, expectedAccessToken.Token, oauthAccessToken.Token)
		})

		t.Run("When request post to altair failed then it will return error", func(t *testing.T) {
			request := entity.GrantTokenRequest{
				RedirectURI:     "http://localhost:8080",
				ClientUID:       "client-uid",
				ClientSecret:    "secret",
				ResponseType:    "token",
				Scopes:          "users",
				ResourceOwnerID: 99,
			}

			networkCfg := mockProvider.NewMockNetworkConfig(mockCtrl)

			encodedJSON, _ := json.Marshal(request)
			network := mockProvider.NewMockNetwork(mockCtrl)
			network.EXPECT().POST(ctx, networkCfg, "/_plugins/oauth/authorizations", bytes.NewBuffer(encodedJSON), gomock.Any(), nil).DoAndReturn(func(ctx context.Context, cfg provider.NetworkConfig, path string, body io.Reader, successBinder interface{}, failedBinder interface{}) error {
				return &entity.ApplicationError{
					Err:        []error{errors.New("internal server error")},
					HTTPStatus: http.StatusInternalServerError,
				}
			})

			grantToken := usecase.GrantToken{}
			_, err := grantToken.Perform(ctx, request, networkCfg, network)

			assert.NotNil(t, err)
		})
	})
}

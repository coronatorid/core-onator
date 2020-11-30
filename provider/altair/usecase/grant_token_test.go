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

			expectedAccessToken := entity.OauthAccessToken{}

			expectedAccessToken.Data.CreatedAt = time.Now()
			expectedAccessToken.Data.ExpiresIn = 3000
			expectedAccessToken.Data.ID = 1
			expectedAccessToken.Data.OauthApplicationID = 1
			expectedAccessToken.Data.RedirectURI = request.RedirectURI
			expectedAccessToken.Data.ResourceOwnerID = request.ResourceOwnerID
			expectedAccessToken.Data.Scopes = request.Scopes
			expectedAccessToken.Data.Token = "some secret token"

			networkCfg := mockProvider.NewMockNetworkConfig(mockCtrl)

			encodedJSON, _ := json.Marshal(request)
			network := mockProvider.NewMockNetwork(mockCtrl)
			network.EXPECT().POST(ctx, networkCfg, "/_plugins/oauth/authorizations", bytes.NewBuffer(encodedJSON), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, cfg provider.NetworkConfig, path string, body io.Reader, successBinder interface{}, failedBinder interface{}) error {
				binder := successBinder.(*entity.OauthAccessToken)
				binder.Data.CreatedAt = expectedAccessToken.Data.CreatedAt
				binder.Data.ExpiresIn = expectedAccessToken.Data.ExpiresIn
				binder.Data.ID = expectedAccessToken.Data.ID
				binder.Data.OauthApplicationID = expectedAccessToken.Data.OauthApplicationID
				binder.Data.RedirectURI = expectedAccessToken.Data.RedirectURI
				binder.Data.ResourceOwnerID = expectedAccessToken.Data.ResourceOwnerID
				binder.Data.Scopes = expectedAccessToken.Data.Scopes
				binder.Data.Token = expectedAccessToken.Data.Token
				return nil
			})

			grantToken := usecase.GrantToken{}
			oauthAccessToken, err := grantToken.Perform(ctx, request, networkCfg, network)

			assert.Nil(t, err)
			assert.Equal(t, expectedAccessToken.Data.CreatedAt, oauthAccessToken.Data.CreatedAt)
			assert.Equal(t, expectedAccessToken.Data.ExpiresIn, oauthAccessToken.Data.ExpiresIn)
			assert.Equal(t, expectedAccessToken.Data.ID, oauthAccessToken.Data.ID)
			assert.Equal(t, expectedAccessToken.Data.OauthApplicationID, oauthAccessToken.Data.OauthApplicationID)
			assert.Equal(t, expectedAccessToken.Data.RedirectURI, oauthAccessToken.Data.RedirectURI)
			assert.Equal(t, expectedAccessToken.Data.ResourceOwnerID, oauthAccessToken.Data.ResourceOwnerID)
			assert.Equal(t, expectedAccessToken.Data.Scopes, oauthAccessToken.Data.Scopes)
			assert.Equal(t, expectedAccessToken.Data.Token, oauthAccessToken.Data.Token)
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
			network.EXPECT().POST(ctx, networkCfg, "/_plugins/oauth/authorizations", bytes.NewBuffer(encodedJSON), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, cfg provider.NetworkConfig, path string, body io.Reader, successBinder interface{}, failedBinder interface{}) error {
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

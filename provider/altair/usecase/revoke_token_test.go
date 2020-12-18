package usecase_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/altair/usecase"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRevokeToken(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := testhelper.NewTestContext()

	t.Run("Perform", func(t *testing.T) {
		t.Run("When request post to altair success then it will return entity.OauthAccessToken", func(t *testing.T) {

			request := entity.RevokeTokenRequest{
				Token: "this-is-some-token",
			}

			networkCfg := mockProvider.NewMockNetworkConfig(mockCtrl)

			encodedJSON, _ := json.Marshal(request)
			network := mockProvider.NewMockNetwork(mockCtrl)
			network.EXPECT().POST(ctx, networkCfg, "/_plugins/oauth/authorizations/revoke", bytes.NewBuffer(encodedJSON), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx provider.Context, cfg provider.NetworkConfig, path string, body io.Reader, successBinder interface{}, failedBinder interface{}) error {
				return nil
			})

			revokeToken := usecase.RevokeToken{}
			err := revokeToken.Perform(ctx, request, networkCfg, network)

			assert.Nil(t, err)
		})

		t.Run("When request post to altair failed then it will return error", func(t *testing.T) {
			request := entity.RevokeTokenRequest{
				Token: "this-is-some-token",
			}

			networkCfg := mockProvider.NewMockNetworkConfig(mockCtrl)

			encodedJSON, _ := json.Marshal(request)
			network := mockProvider.NewMockNetwork(mockCtrl)
			network.EXPECT().POST(ctx, networkCfg, "/_plugins/oauth/authorizations/revoke", bytes.NewBuffer(encodedJSON), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx provider.Context, cfg provider.NetworkConfig, path string, body io.Reader, successBinder interface{}, failedBinder interface{}) error {
				return &entity.ApplicationError{
					Err:        []error{errors.New("internal server error")},
					HTTPStatus: http.StatusInternalServerError,
				}
			})

			revokeToken := usecase.RevokeToken{}
			err := revokeToken.Perform(ctx, request, networkCfg, network)

			assert.NotNil(t, err)
		})
	})
}

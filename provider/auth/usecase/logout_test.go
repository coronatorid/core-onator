package usecase_test

import (
	"testing"

	"github.com/coronatorid/core-onator/entity"
	"github.com/stretchr/testify/assert"

	"github.com/coronatorid/core-onator/provider/auth/usecase"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/golang/mock/gomock"
)

func TestLogout(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := testhelper.NewTestContext()
	altair := mockProvider.NewMockAltair(mockCtrl)
	request := entity.RevokeTokenRequest{
		Token: "this-token-is-secret",
	}

	t.Run("Perform", func(t *testing.T) {
		t.Run("When request success it will return nil", func(t *testing.T) {
			revokeToken := usecase.Logout{}
			altair.EXPECT().RevokeToken(ctx, request).Return(nil)
			assert.Nil(t, revokeToken.Perform(ctx, request, altair))
		})
	})
}

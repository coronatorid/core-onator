package usecase_test

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/provider/user/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindByPhoneNumber(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	precompiledPhoneNumber := "+6289787657891"

	// We should protect our user no matter what
	h := sha256.New()
	_, _ = h.Write([]byte(fmt.Sprintf("%sXXX%s", precompiledPhoneNumber, os.Getenv("APP_ENCRIPTION_KEY"))))
	phoneNumber := fmt.Sprintf("%x", h.Sum(nil))

	t.Run("Perform", func(t *testing.T) {
		t.Run("When user is found then it will return user", func(t *testing.T) {
			row := mockProvider.NewMockRow(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().QueryRowContext(ctx, "find-user", "select id, phone, state, created_at, updated_at from users where phone = ?", phoneNumber).Return(row)
			row.EXPECT().Scan(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			find := &usecase.FindByPhoneNumber{}
			_, err := find.Perform(ctx, precompiledPhoneNumber, db)
			assert.Nil(t, err)
		})

		t.Run("When user is not found then it will return error", func(t *testing.T) {
			row := mockProvider.NewMockRow(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().QueryRowContext(ctx, "find-user", "select id, phone, state, created_at, updated_at from users where phone = ?", phoneNumber).Return(row)
			row.EXPECT().Scan(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(provider.ErrDBNotFound)

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("user not found")},
				HTTPStatus: http.StatusNotFound,
			}

			find := &usecase.FindByPhoneNumber{}
			_, err := find.Perform(ctx, precompiledPhoneNumber, db)
			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
		})

		t.Run("When there is unexpected error in database then it will return error", func(t *testing.T) {
			row := mockProvider.NewMockRow(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().QueryRowContext(ctx, "find-user", "select id, phone, state, created_at, updated_at from users where phone = ?", phoneNumber).Return(row)
			row.EXPECT().Scan(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("unexpected error"))

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			}

			find := &usecase.FindByPhoneNumber{}
			_, err := find.Perform(ctx, precompiledPhoneNumber, db)
			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
		})
	})
}

package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/provider/user/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	ID := 1

	t.Run("Perform", func(t *testing.T) {
		t.Run("When user is found then it will return user", func(t *testing.T) {
			row := mockProvider.NewMockRow(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().QueryRowContext(ctx, "find-user", "select id, phone, state, created_at, updated_at from users where id = ?", ID).Return(row)
			row.EXPECT().Scan(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			find := &usecase.Find{}
			_, err := find.Perform(ctx, ID, db)
			assert.Nil(t, err)
		})

		t.Run("When user is not found then it will return error", func(t *testing.T) {
			row := mockProvider.NewMockRow(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().QueryRowContext(ctx, "find-user", "select id, phone, state, created_at, updated_at from users where id = ?", ID).Return(row)
			row.EXPECT().Scan(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(provider.ErrDBNotFound)

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("user not found")},
				HTTPStatus: http.StatusNotFound,
			}

			find := &usecase.Find{}
			_, err := find.Perform(ctx, ID, db)
			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
		})

		t.Run("When there is unexpected error in database then it will return error", func(t *testing.T) {
			row := mockProvider.NewMockRow(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().QueryRowContext(ctx, "find-user", "select id, phone, state, created_at, updated_at from users where id = ?", ID).Return(row)
			row.EXPECT().Scan(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("unexpected error"))

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			}

			find := &usecase.Find{}
			_, err := find.Perform(ctx, ID, db)
			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
		})
	})
}
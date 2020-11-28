package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/coronatorid/core-onator/entity"
	"github.com/stretchr/testify/assert"

	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/provider/user/usecase"

	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	phoneNumber := "+6289765430918"

	t.Run("Perform", func(t *testing.T) {
		t.Run("When user creation is success then it will return newly created id", func(t *testing.T) {
			userInsertable := entity.UserInsertable{
				PhoneNumber: phoneNumber,
			}

			result := mockProvider.NewMockResult(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().ExecContext(ctx, "user-create", "insert into users (phone, state, created_at, updated_at) values(?, 1, now(), now())", userInsertable.PhoneNumber).Return(result, nil)
			result.EXPECT().LastInsertId().Return(int64(99), nil)

			create := &usecase.Create{}
			ID, err := create.Perform(ctx, userInsertable, db)

			assert.Nil(t, err)
			assert.Equal(t, 99, ID)
		})

		t.Run("When user creation failed because of database execution it will return error", func(t *testing.T) {
			userInsertable := entity.UserInsertable{
				PhoneNumber: phoneNumber,
			}

			db := mockProvider.NewMockDB(mockCtrl)
			db.EXPECT().ExecContext(ctx, "user-create", "insert into users (phone, state, created_at, updated_at) values(?, 1, now(), now())", userInsertable.PhoneNumber).Return(nil, errors.New("unexpected error"))

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			}

			create := &usecase.Create{}
			ID, err := create.Perform(ctx, userInsertable, db)

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Equal(t, 0, ID)
		})

		t.Run("When user creation failed because of last inserted id", func(t *testing.T) {
			userInsertable := entity.UserInsertable{
				PhoneNumber: phoneNumber,
			}

			result := mockProvider.NewMockResult(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().ExecContext(ctx, "user-create", "insert into users (phone, state, created_at, updated_at) values(?, 1, now(), now())", userInsertable.PhoneNumber).Return(result, nil)
			result.EXPECT().LastInsertId().Return(int64(0), errors.New("unexpected error"))

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("internal server error when acquiring last inserted id")},
				HTTPStatus: http.StatusInternalServerError,
			}

			create := &usecase.Create{}
			ID, err := create.Perform(ctx, userInsertable, db)

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Equal(t, 0, ID)
		})
	})
}

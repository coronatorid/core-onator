package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/coronatorid/core-onator/entity"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/provider/tracker/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	userID := 99
	lat := 21.3875022
	long := 39.8628782

	t.Run("Perform", func(t *testing.T) {
		t.Run("When location creation is success then it will return newly created id", func(t *testing.T) {
			locationInsertable := entity.LocationInsertable{
				UserID: userID,
				Lat:    lat,
				Long:   long,
			}

			result := mockProvider.NewMockResult(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().ExecContext(ctx, "location-create", "insert into locations (user_id, lat, long, created_at, updated_at) values(?, ?, ?, now(), now())", locationInsertable.UserID, locationInsertable.Lat, locationInsertable.Long).Return(result, nil)
			result.EXPECT().LastInsertId().Return(int64(99), nil)

			create := &usecase.Create{}
			ID, err := create.Perform(ctx, locationInsertable, db)

			assert.Nil(t, err)
			assert.Equal(t, 99, ID)
		})

		t.Run("When location creation failed because of database execution it will return error", func(t *testing.T) {
			locationInsertable := entity.LocationInsertable{
				UserID: userID,
				Lat:    lat,
				Long:   long,
			}

			db := mockProvider.NewMockDB(mockCtrl)
			db.EXPECT().ExecContext(ctx, "location-create", "insert into locations (user_id, lat, long, created_at, updated_at) values(?, ?, ?, now(), now())", locationInsertable.UserID, locationInsertable.Lat, locationInsertable.Long).Return(nil, errors.New("unexpected error"))

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			}

			create := &usecase.Create{}
			ID, err := create.Perform(ctx, locationInsertable, db)

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Equal(t, 0, ID)
		})

		t.Run("When location creation failed because of last inserted id", func(t *testing.T) {
			locationInsertable := entity.LocationInsertable{
				UserID: userID,
				Lat:    lat,
				Long:   long,
			}

			result := mockProvider.NewMockResult(mockCtrl)
			db := mockProvider.NewMockDB(mockCtrl)

			db.EXPECT().ExecContext(ctx, "location-create", "insert into locations (user_id, lat, long, created_at, updated_at) values(?, ?, ?, now(), now())", locationInsertable.UserID, locationInsertable.Lat, locationInsertable.Long).Return(result, nil)
			result.EXPECT().LastInsertId().Return(int64(0), errors.New("unexpected error"))

			expectedError := &entity.ApplicationError{
				Err:        []error{errors.New("internal server error when acquiring last inserted id")},
				HTTPStatus: http.StatusInternalServerError,
			}

			create := &usecase.Create{}
			ID, err := create.Perform(ctx, locationInsertable, db)

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Equal(t, 0, ID)
		})
	})
}

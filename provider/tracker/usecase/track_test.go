package usecase_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/stretchr/testify/assert"

	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/provider/tracker/usecase"

	"github.com/golang/mock/gomock"
)

func TestTrack(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := testhelper.NewTestContext()

	locationID := 1
	userID := 999
	lat := 21.3875022
	long := 39.8628782

	t.Run("Perform", func(t *testing.T) {
		t.Run("When track process success it will return entity.Location", func(t *testing.T) {
			request := entity.TrackRequest{
				Lat:  lat,
				Long: long,
			}

			expectedLocation := entity.Location{
				ID:        locationID,
				UserID:    userID,
				Lat:       lat,
				Long:      long,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}

			tracker := mockProvider.NewMockTracker(mockCtrl)
			tracker.EXPECT().Create(ctx, entity.LocationInsertable{
				UserID: userID,
				Lat:    lat,
				Long:   long,
			}).Return(locationID, nil)
			tracker.EXPECT().Find(ctx, locationID).Return(expectedLocation, nil)

			track := &usecase.Track{}
			location, err := track.Perform(ctx, userID, request, tracker)

			assert.Nil(t, err)
			assert.Equal(t, expectedLocation.ID, location.ID)
			assert.Equal(t, expectedLocation.UserID, location.UserID)
			assert.Equal(t, expectedLocation.Lat, location.Lat)
			assert.Equal(t, expectedLocation.Long, location.Long)
			assert.Equal(t, expectedLocation.CreatedAt, location.CreatedAt)
			assert.Equal(t, expectedLocation.UpdatedAt, location.UpdatedAt)
		})

		t.Run("When create location failed it will return error", func(t *testing.T) {
			request := entity.TrackRequest{
				Lat:  lat,
				Long: long,
			}

			tracker := mockProvider.NewMockTracker(mockCtrl)
			tracker.EXPECT().Create(ctx, entity.LocationInsertable{
				UserID: userID,
				Lat:    lat,
				Long:   long,
			}).Return(0, &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			})

			track := &usecase.Track{}
			_, err := track.Perform(ctx, userID, request, tracker)

			assert.NotNil(t, err)
		})

		t.Run("When find location failed it will return error", func(t *testing.T) {
			request := entity.TrackRequest{
				Lat:  lat,
				Long: long,
			}

			tracker := mockProvider.NewMockTracker(mockCtrl)
			tracker.EXPECT().Create(ctx, entity.LocationInsertable{
				UserID: userID,
				Lat:    lat,
				Long:   long,
			}).Return(locationID, nil)
			tracker.EXPECT().Find(ctx, locationID).Return(entity.Location{}, &entity.ApplicationError{
				Err:        []error{errors.New("location not found")},
				HTTPStatus: http.StatusNotFound,
			})

			track := &usecase.Track{}
			_, err := track.Perform(ctx, userID, request, tracker)

			assert.NotNil(t, err)
		})
	})
}

package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/entity"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"

	"github.com/coronatorid/core-onator/provider/user/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrFindUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	phoneNumber := "+6289787657281"

	expectedUser := entity.User{
		ID:        99,
		Phone:     phoneNumber,
		State:     1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("Perform", func(t *testing.T) {
		t.Run("When user is not found then it will create a new one", func(t *testing.T) {
			userProvider := mockProvider.NewMockUser(mockCtrl)
			userProvider.EXPECT().FindByPhoneNumber(ctx, phoneNumber).Return(entity.User{}, &entity.ApplicationError{
				Err:        []error{errors.New("user not found")},
				HTTPStatus: http.StatusNotFound,
			})
			userProvider.EXPECT().Create(ctx, entity.UserInsertable{PhoneNumber: phoneNumber}).Return(expectedUser.ID, nil)
			userProvider.EXPECT().Find(ctx, expectedUser.ID).Return(expectedUser, nil)

			createOrFindUser := usecase.CreateOrFindUser{}
			user, err := createOrFindUser.Perform(ctx, phoneNumber, userProvider)
			assert.Nil(t, err)
			assert.Equal(t, expectedUser.ID, user.ID)
			assert.Equal(t, expectedUser.Phone, user.Phone)
			assert.Equal(t, expectedUser.State, user.State)
			assert.Equal(t, expectedUser.CreatedAt, user.CreatedAt)
			assert.Equal(t, expectedUser.UpdatedAt, user.UpdatedAt)
		})

		t.Run("When user is found then it will return old the old one", func(t *testing.T) {
			userProvider := mockProvider.NewMockUser(mockCtrl)
			userProvider.EXPECT().FindByPhoneNumber(ctx, phoneNumber).Return(expectedUser, nil)

			createOrFindUser := usecase.CreateOrFindUser{}
			user, err := createOrFindUser.Perform(ctx, phoneNumber, userProvider)
			assert.Nil(t, err)
			assert.Equal(t, expectedUser.ID, user.ID)
			assert.Equal(t, expectedUser.Phone, user.Phone)
			assert.Equal(t, expectedUser.State, user.State)
			assert.Equal(t, expectedUser.CreatedAt, user.CreatedAt)
			assert.Equal(t, expectedUser.UpdatedAt, user.UpdatedAt)
		})

		t.Run("When there is database error then it will return error", func(t *testing.T) {
			userProvider := mockProvider.NewMockUser(mockCtrl)
			userProvider.EXPECT().FindByPhoneNumber(ctx, phoneNumber).Return(entity.User{}, &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			})

			createOrFindUser := usecase.CreateOrFindUser{}
			_, err := createOrFindUser.Perform(ctx, phoneNumber, userProvider)
			assert.NotNil(t, err)
		})

		t.Run("When there is database error when creating user then it will return error", func(t *testing.T) {
			userProvider := mockProvider.NewMockUser(mockCtrl)
			userProvider.EXPECT().FindByPhoneNumber(ctx, phoneNumber).Return(entity.User{}, &entity.ApplicationError{
				Err:        []error{errors.New("user not found")},
				HTTPStatus: http.StatusNotFound,
			})
			userProvider.EXPECT().Create(ctx, entity.UserInsertable{PhoneNumber: phoneNumber}).Return(0, &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			})

			createOrFindUser := usecase.CreateOrFindUser{}
			_, err := createOrFindUser.Perform(ctx, phoneNumber, userProvider)
			assert.NotNil(t, err)
		})

		t.Run("When there is database error after creating user then it will return error", func(t *testing.T) {
			userProvider := mockProvider.NewMockUser(mockCtrl)
			userProvider.EXPECT().FindByPhoneNumber(ctx, phoneNumber).Return(entity.User{}, &entity.ApplicationError{
				Err:        []error{errors.New("user not found")},
				HTTPStatus: http.StatusNotFound,
			})
			userProvider.EXPECT().Create(ctx, entity.UserInsertable{PhoneNumber: phoneNumber}).Return(expectedUser.ID, nil)
			userProvider.EXPECT().Find(ctx, expectedUser.ID).Return(expectedUser, &entity.ApplicationError{
				Err:        []error{errors.New("user not found")},
				HTTPStatus: http.StatusNotFound,
			})

			createOrFindUser := usecase.CreateOrFindUser{}
			_, err := createOrFindUser.Perform(ctx, phoneNumber, userProvider)
			assert.NotNil(t, err)
		})
	})
}

package usecase_test

import (
	"errors"
	"testing"

	"github.com/coronatorid/core-onator/entity"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/provider/report/usecase"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/coronatorid/core-onator/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := testhelper.NewTestContext()

	t.Run("Perform", func(t *testing.T) {
		t.Run("When reported cases creation is success then it will return newly created id", func(t *testing.T) {
			insertable := entity.ReportInsertable{
				UserID:    10,
				ImagePath: "/opt/api/storage/xxx.png",
			}

			result := mockProvider.NewMockResult(mockCtrl)
			tx := mockProvider.NewMockTX(mockCtrl)

			tx.EXPECT().ExecContext(ctx, "reported-cases-create", "insert into users (user_id, status, image_path, image_deleted, created_at, updated_at) values(?, 2, ?, 0, now(), now())", insertable.UserID, insertable.ImagePath).Return(result, nil)
			result.EXPECT().LastInsertId().Return(int64(99), nil)

			createReportCases := &usecase.CreateReportCases{}
			ID, err := createReportCases.Perform(ctx, insertable, tx)

			assert.Nil(t, err)
			assert.Equal(t, 99, ID)
		})

		t.Run("When user creation failed because of database execution it will return error", func(t *testing.T) {
			insertable := entity.ReportInsertable{
				UserID:    10,
				ImagePath: "/opt/api/storage/xxx.png",
			}

			tx := mockProvider.NewMockTX(mockCtrl)
			tx.EXPECT().ExecContext(ctx, "reported-cases-create", "insert into users (user_id, status, image_path, image_deleted, created_at, updated_at) values(?, 2, ?, 0, now(), now())", insertable.UserID, insertable.ImagePath).Return(nil, errors.New("unexpected error"))

			expectedError := util.CreateInternalServerError(ctx)

			createReportCases := &usecase.CreateReportCases{}
			ID, err := createReportCases.Perform(ctx, insertable, tx)

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Equal(t, 0, ID)
		})

		t.Run("When user creation failed because of last inserted id", func(t *testing.T) {
			insertable := entity.ReportInsertable{
				UserID:    10,
				ImagePath: "/opt/api/storage/xxx.png",
			}

			result := mockProvider.NewMockResult(mockCtrl)
			tx := mockProvider.NewMockTX(mockCtrl)

			tx.EXPECT().ExecContext(ctx, "reported-cases-create", "insert into users (user_id, status, image_path, image_deleted, created_at, updated_at) values(?, 2, ?, 0, now(), now())", insertable.UserID, insertable.ImagePath).Return(result, nil)
			result.EXPECT().LastInsertId().Return(int64(0), errors.New("unexpected error"))

			expectedError := util.CreateInternalServerError(ctx)

			createReportCases := &usecase.CreateReportCases{}
			ID, err := createReportCases.Perform(ctx, insertable, tx)

			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Equal(t, expectedError.HTTPStatus, err.HTTPStatus)
			assert.Equal(t, 0, ID)
		})
	})
}

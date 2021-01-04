package report_test

import (
	"testing"

	"github.com/coronatorid/core-onator/entity"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/provider/report"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := testhelper.NewTestContext()
	db := mockProvider.NewMockDB(mockCtrl)
	reportProvider := report.Fabricate(db)
	t.Run("CreateReportCases", func(t *testing.T) {
		insertable := entity.ReportInsertable{
			UserID:    10,
			ImagePath: "/opt/api/storage/xxx.png",
		}

		result := mockProvider.NewMockResult(mockCtrl)
		tx := mockProvider.NewMockTX(mockCtrl)

		tx.EXPECT().ExecContext(ctx, "reported-cases-create", "insert into users (user_id, status, image_path, image_deleted, created_at, updated_at) values(?, 2, ?, 0, now(), now())", insertable.UserID, insertable.ImagePath).Return(result, nil)
		result.EXPECT().LastInsertId().Return(int64(99), nil)

		ID, err := reportProvider.CreateReportCases(ctx, insertable, tx)

		assert.Nil(t, err)
		assert.Equal(t, 99, ID)
	})
}
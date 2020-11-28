package adapter_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/infrastructure/adapter"
	"github.com/stretchr/testify/assert"

	mockProvider "github.com/coronatorid/core-onator/provider/mocks"

	"github.com/golang/mock/gomock"
)

func TestSQLRow(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("Scan", func(t *testing.T) {
		t.Run("When there is no error then it will return nil", func(t *testing.T) {
			row := mockProvider.NewMockRow(mockCtrl)
			row.EXPECT().Scan(gomock.Any()).Return(nil)

			var x int

			sqlRow := adapter.AdaptSQLRow(row)
			assert.Nil(t, sqlRow.Scan(&x))
		})

		t.Run("When there is error then it will return error", func(t *testing.T) {
			row := mockProvider.NewMockRow(mockCtrl)
			row.EXPECT().Scan(gomock.Any()).Return(errors.New("unexpected error"))

			var x int

			sqlRow := adapter.AdaptSQLRow(row)
			assert.NotNil(t, sqlRow.Scan(&x))
		})

		t.Run("When there is sql.ErrNoRows then it will return provider.ErrDBNotFound", func(t *testing.T) {
			row := mockProvider.NewMockRow(mockCtrl)
			row.EXPECT().Scan(gomock.Any()).Return(sql.ErrNoRows)

			var x int

			sqlRow := adapter.AdaptSQLRow(row)
			assert.Equal(t, provider.ErrDBNotFound, sqlRow.Scan(&x))
		})
	})
}

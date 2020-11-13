package infrastructure_test

import (
	"testing"

	mockProvider "github.com/coronatorid/core-onator/provider/mocks"

	"github.com/coronatorid/core-onator/provider/infrastructure"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInfrastructure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("MYSQL", func(t *testing.T) {
		t.Run("Can connect to mysql", func(t *testing.T) {
			i := infrastructure.Fabricate()
			db, err := i.MYSQL()
			assert.Nil(t, err)
			assert.NotNil(t, db)
		})

		t.Run("Already have connection established to mysql", func(t *testing.T) {
			i := infrastructure.Fabricate()

			db, err := i.MYSQL()
			assert.Nil(t, err)
			assert.NotNil(t, db)

			db, err = i.MYSQL()
			assert.Nil(t, err)
			assert.NotNil(t, db)
		})
	})

	t.Run("InjectCommand", func(t *testing.T) {
		t.Run("Given provider.Command", func(t *testing.T) {
			i := infrastructure.Fabricate()

			commandProvider := mockProvider.NewMockCommand(mockCtrl)
			commandProvider.EXPECT().InjectCommand(gomock.Any()).Times(1)

			_ = i.FabricateCommand(commandProvider)
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("When mysql connection is not nil it close mysql connection", func(t *testing.T) {
			i := infrastructure.Fabricate()
			_, _ = i.MYSQL()
			i.Close()
		})
	})

	t.Run("Memcached", func(t *testing.T) {
		t.Run("Return memcached object", func(t *testing.T) {
			i := infrastructure.Fabricate()
			assert.NotPanics(t, func() {
				_ = i.Memcached()
			})
		})
	})
}

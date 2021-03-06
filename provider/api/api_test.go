package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/provider/api"
	"github.com/coronatorid/core-onator/provider/api/handler"
	"github.com/coronatorid/core-onator/provider/inappcron"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/golang/mock/gomock"
)

func TestAPI(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inappcron := inappcron.Fabricate()
	t.Run("Run and Shutdown", func(t *testing.T) {
		apiEngine := api.Fabricate(inappcron)

		go func() {
			apiEngine.Run()
		}()

		time.Sleep(time.Millisecond * 500)

		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
		defer cancelFunc()

		_ = apiEngine.Shutdown(ctx)
	})

	t.Run("FabricateCommand", func(t *testing.T) {
		t.Run("Given provider.Command", func(t *testing.T) {
			apiEngine := api.Fabricate(inappcron)

			commandProvider := mockProvider.NewMockCommand(mockCtrl)
			commandProvider.EXPECT().InjectCommand(gomock.Any()).Times(1)

			apiEngine.FabricateCommand(commandProvider)
		})
	})

	t.Run("Inject API", func(t *testing.T) {
		apiEngine := api.Fabricate(inappcron)
		apiEngine.InjectAPI(handler.NewHealth())
	})

	t.Run("Normal API Scenario", func(t *testing.T) {

	})
}

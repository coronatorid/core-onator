package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/provider/api"
	"github.com/coronatorid/core-onator/provider/api/handler"
	"github.com/golang/mock/gomock"
)

func TestAPI(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("Run and Shutdown", func(t *testing.T) {
		apiEngine := api.Fabricate()

		go func() {
			apiEngine.Run()
		}()

		time.Sleep(time.Millisecond * 500)

		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
		defer cancelFunc()

		_ = apiEngine.Shutdown(ctx)
	})

	t.Run("Inject API", func(t *testing.T) {
		apiEngine := api.Fabricate()
		apiEngine.InjectAPI(handler.NewHealth())
	})

	t.Run("Normal API Scenario", func(t *testing.T) {

	})
}

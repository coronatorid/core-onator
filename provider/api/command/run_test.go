package command_test

import (
	"syscall"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/provider/api"
	"github.com/coronatorid/core-onator/provider/api/command"
	"github.com/coronatorid/core-onator/provider/inappcron"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	apiEngine := api.Fabricate(inappcron.Fabricate())
	runAPICMd := command.NewRun(apiEngine, inappcron.Fabricate())

	t.Run("Use", func(t *testing.T) {
		assert.Equal(t, "run:api", runAPICMd.Use())
	})

	t.Run("Example", func(t *testing.T) {
		assert.Equal(t, "run:api", runAPICMd.Example())
	})

	t.Run("Short", func(t *testing.T) {
		assert.Equal(t, "Run API engine", runAPICMd.Short())
	})

	t.Run("Run", func(t *testing.T) {
		t.Run("Given args", func(t *testing.T) {
			go func() {
				<-time.After(time.Millisecond * 500)
				syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
			}()

			assert.NotPanics(t, func() {
				runAPICMd.Run([]string{})
			})
		})
	})
}

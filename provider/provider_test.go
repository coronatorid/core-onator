package provider_test

import (
	"testing"

	mockCommand "github.com/coronatorid/core-onator/provider/command/mocks"

	"github.com/coronatorid/core-onator/provider"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	p := provider.Fabricate()

	t.Run("InjectCommand", func(t *testing.T) {
		commandScaffolding := mockCommand.NewMockScaffold(mockCtrl)
		commandScaffolding.EXPECT().Use().Return("others")
		commandScaffolding.EXPECT().Short().Return("Others command")
		commandScaffolding.EXPECT().Example().Return("others [command]")
		assert.NotPanics(t, func() { p.InjectCommand(commandScaffolding) })
	})

	t.Run("Infrastructure", func(t *testing.T) {
		assert.NotPanics(t, func() { p.Infrastructure() })
	})
}

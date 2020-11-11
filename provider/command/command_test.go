package command_test

import (
	"testing"

	mockProvider "github.com/coronatorid/core-onator/provider/mocks"

	"github.com/golang/mock/gomock"

	"github.com/coronatorid/core-onator/provider/command"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("coronator", func(t *testing.T) {
		cmd := command.Fabricate()
		cmd.SetArgs([]string{"coronator"})
		assert.Nil(t, cmd.Execute())
	})

	t.Run("others", func(t *testing.T) {
		cmd := command.Fabricate()
		cmd.SetArgs([]string{"others"})

		scaffolding := mockProvider.NewMockCommandScaffold(mockCtrl)

		scaffolding.EXPECT().Use().Return("others")
		scaffolding.EXPECT().Short().Return("Others command")
		scaffolding.EXPECT().Example().Return("others [command]")
		scaffolding.EXPECT().Run([]string{})
		cmd.InjectCommand(scaffolding)
		assert.Nil(t, cmd.Execute())
	})
}

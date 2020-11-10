package command_test

import (
	"testing"

	"github.com/coronatorid/core-onator/provider/command"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {

	t.Run("coronator", func(t *testing.T) {
		cmd := command.Fabricate()
		cmd.SetArgs([]string{"coronator"})
		assert.Nil(t, cmd.Execute())
	})
}

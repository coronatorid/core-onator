package provider_test

import (
	"testing"

	"github.com/coronatorid/core-onator/provider"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	p := provider.Fabricate()

	t.Run("Command", func(t *testing.T) {
		assert.NotPanics(t, func() { p.Command() })
	})

	t.Run("Infrastructure", func(t *testing.T) {
		assert.NotPanics(t, func() { p.Infrastructure() })
	})
}

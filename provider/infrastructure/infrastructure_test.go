package infrastructure_test

import (
	"testing"

	"github.com/coronatorid/core-onator/provider/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestInfrastructure(t *testing.T) {
	t.Run("MYSQL", func(t *testing.T) {
		t.Run("Can connect to mysql", func(t *testing.T) {
			i := infrastructure.Fabricate()
			_, err := i.MYSQL()
			assert.Nil(t, err)
		})

		t.Run("Already have connection established to mysql", func(t *testing.T) {
			i := infrastructure.Fabricate()

			_, err := i.MYSQL()
			assert.Nil(t, err)

			_, err = i.MYSQL()
			assert.Nil(t, err)
		})
	})
}

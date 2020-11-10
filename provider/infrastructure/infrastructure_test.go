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
}

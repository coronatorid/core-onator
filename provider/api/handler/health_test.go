package handler_test

import (
	"testing"

	mockProvider "github.com/coronatorid/core-onator/provider/mocks"

	"github.com/coronatorid/core-onator/provider/api/handler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	healthHandler := handler.NewHealth()

	t.Run("Path", func(t *testing.T) {
		assert.Equal(t, "/health", healthHandler.Path())
	})

	t.Run("Method", func(t *testing.T) {
		assert.Equal(t, "GET", healthHandler.Method())
	})

	t.Run("Handle", func(t *testing.T) {
		apiContext := mockProvider.NewMockAPIContext(mockCtrl)
		assert.NotPanics(t, func() {
			healthHandler.Handle(apiContext)
		})
	})
}

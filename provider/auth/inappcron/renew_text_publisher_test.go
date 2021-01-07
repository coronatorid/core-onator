package inappcron_test

import (
	"testing"
	"time"

	"github.com/coronatorid/core-onator/provider/auth/inappcron"
	mockProvider "github.com/coronatorid/core-onator/provider/mocks"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestRenewTextPublisher(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("Delay", func(t *testing.T) {
		auth := mockProvider.NewMockAuth(mockCtrl)
		assert.Equal(t, 15*time.Minute, inappcron.NewRenewTextPublisher(auth, testhelper.WhatsappPublisher{Controller: mockCtrl}.NewWhatsappPublisher).Delay())
	})

	t.Run("Close", func(t *testing.T) {
		assert.NotPanics(t, func() {
			auth := mockProvider.NewMockAuth(mockCtrl)
			inappcron.NewRenewTextPublisher(auth, testhelper.WhatsappPublisher{Controller: mockCtrl}.NewWhatsappPublisher).Close()
		})
	})

	t.Run("Run", func(t *testing.T) {
		t.Run("When success it will return nil", func(t *testing.T) {
			auth := mockProvider.NewMockAuth(mockCtrl)
			auth.EXPECT().RenewTextPublisher(gomock.Any()).Times(1)
			assert.Nil(t, inappcron.NewRenewTextPublisher(auth, testhelper.WhatsappPublisher{Controller: mockCtrl}.NewWhatsappPublisher).Run())
		})

		t.Run("When error it will return error", func(t *testing.T) {
			auth := mockProvider.NewMockAuth(mockCtrl)
			assert.NotNil(t, inappcron.NewRenewTextPublisher(auth, testhelper.WhatsappPublisher{Controller: mockCtrl}.NewWhatsappPublisherError).Run())
		})
	})

	t.Run("Name", func(t *testing.T) {
		auth := mockProvider.NewMockAuth(mockCtrl)
		assert.Equal(t, "auth/renew_text_publisher", inappcron.NewRenewTextPublisher(auth, testhelper.WhatsappPublisher{Controller: mockCtrl}.NewWhatsappPublisher).Name())
	})
}

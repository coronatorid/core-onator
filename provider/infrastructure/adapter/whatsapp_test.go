package adapter_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/Rhymen/go-whatsapp"
	"github.com/coronatorid/core-onator/provider/infrastructure/adapter"
	mockAdapter "github.com/coronatorid/core-onator/provider/infrastructure/adapter/mocks"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestWhatsapp(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := testhelper.NewTestContext()
	phoneNumber := "+6287687901240"
	message := "OTP: 130419"

	t.Run("Publish", func(t *testing.T) {
		t.Run("When whatsapp client successfully publish a message then it will return nil", func(t *testing.T) {
			whatsappClient := mockAdapter.NewMockWhatsappClient(mockCtrl)
			whatsappClient.EXPECT().Send(gomock.Any()).DoAndReturn(func(msg interface{}) (string, error) {
				assertedMessage, ok := msg.(whatsapp.TextMessage)

				assert.True(t, ok)
				assert.Equal(t, fmt.Sprintf("%s@s.whatsapp.net", strings.ReplaceAll(phoneNumber, "+", "")), assertedMessage.Info.RemoteJid)
				assert.Equal(t, message, assertedMessage.Text)

				return "OK", nil
			})

			textPublisher := adapter.AdaptWhatsapp(whatsappClient)
			err := textPublisher.Publish(ctx, phoneNumber, message)
			assert.Nil(t, err)
		})

		t.Run("When whatsapp client failed to publish a message then it will return error", func(t *testing.T) {
			whatsappClient := mockAdapter.NewMockWhatsappClient(mockCtrl)
			whatsappClient.EXPECT().Send(gomock.Any()).DoAndReturn(func(msg interface{}) (string, error) {
				assertedMessage, ok := msg.(whatsapp.TextMessage)

				assert.True(t, ok)
				assert.Equal(t, fmt.Sprintf("%s@s.whatsapp.net", strings.ReplaceAll(phoneNumber, "+", "")), assertedMessage.Info.RemoteJid)
				assert.Equal(t, message, assertedMessage.Text)

				return "ERROR", errors.New("unexpected error")
			})

			textPublisher := adapter.AdaptWhatsapp(whatsappClient)
			err := textPublisher.Publish(ctx, phoneNumber, message)
			assert.NotNil(t, err)
		})
	})
}

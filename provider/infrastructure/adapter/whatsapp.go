package adapter

import (
	"fmt"
	"strings"

	"github.com/Rhymen/go-whatsapp"
	"github.com/coronatorid/core-onator/provider"
)

// Whatsapp wrap whatsapp connection into TextPublisher interface
type Whatsapp struct {
	client WhatsappClient
}

//go:generate mockgen -source=./whatsapp.go -destination=./mocks/whatsapp_mock.go -package mockAdapter

// WhatsappClient wrap default whatsapp client
type WhatsappClient interface {
	Send(msg interface{}) (string, error)
}

// AdaptWhatsapp adapt TextPublisher interface
func AdaptWhatsapp(client WhatsappClient) *Whatsapp {
	return &Whatsapp{client: client}
}

// Publish message to whatsapp
func (w *Whatsapp) Publish(ctx provider.Context, phoneNumber, message string) error {
	_, err := w.client.Send(whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			FromMe:    true,
			RemoteJid: fmt.Sprintf("%s@s.whatsapp.net", strings.ReplaceAll(phoneNumber, "+", "")),
		},
		Text: message,
	})
	return err
}

package testhelper

import (
	"errors"

	mockProvider "github.com/coronatorid/core-onator/provider/mocks"

	"github.com/golang/mock/gomock"

	"github.com/coronatorid/core-onator/provider"
)

type WhatsappPublisher struct {
	Controller *gomock.Controller
}

func (w WhatsappPublisher) NewWhatsappPublisher() (provider.TextPublisher, error) {
	return mockProvider.NewMockTextPublisher(w.Controller), nil
}

func (WhatsappPublisher) NewWhatsappPublisherError() (provider.TextPublisher, error) {
	return nil, errors.New("Unexpected error")
}

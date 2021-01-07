package inappcron

import (
	"time"

	"github.com/coronatorid/core-onator/provider"
)

// RenewTextPublisher is in app cronjob to renew whatsapp connection
type RenewTextPublisher struct {
	publisherFabricator func() (provider.TextPublisher, error)
	auth                provider.Auth
}

// NewRenewTextPublisher ...
func NewRenewTextPublisher(auth provider.Auth, publisherFabricator func() (provider.TextPublisher, error)) *RenewTextPublisher {
	return &RenewTextPublisher{
		auth:                auth,
		publisherFabricator: publisherFabricator,
	}
}

// Delay between each run
func (r *RenewTextPublisher) Delay() time.Duration {
	return 15 * time.Minute
}

// Close inapp cronjob
func (r *RenewTextPublisher) Close() {
	// Not doing anything
}

// Run inapp cronjob
func (r *RenewTextPublisher) Run() error {
	textPublisher, err := r.publisherFabricator()
	if err != nil {
		return err
	}

	r.auth.RenewTextPublisher(textPublisher)

	return nil
}

// Name of in application cronjob
func (r *RenewTextPublisher) Name() string {
	return "auth/renew_text_publisher"
}

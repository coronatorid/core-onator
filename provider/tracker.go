package provider

import (
	"context"

	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./tracker.go -destination=./mocks/tracker_mock.go -package mockProvider

// Tracker logic of coronator
type Tracker interface {
	Track(ctx context.Context, request entity.TrackRequest) *entity.ApplicationError
}

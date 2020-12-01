package provider

import (
	"context"

	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./tracker.go -destination=./mocks/tracker_mock.go -package mockProvider

// Tracker logic of coronator
type Tracker interface {
	Track(ctx context.Context, userID int, request entity.TrackRequest) (entity.Location, *entity.ApplicationError)
	Create(ctx context.Context, locationInsertable entity.LocationInsertable) (int, *entity.ApplicationError)
	Find(ctx context.Context, locationID int) (entity.Location, *entity.ApplicationError)
}

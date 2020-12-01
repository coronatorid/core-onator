package tracker

import (
	"context"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/tracker/usecase"
)

// Tracker ...
type Tracker struct {
	db provider.DB
}

// Fabricate tracker
func Fabricate(db provider.DB) *Tracker {
	return &Tracker{db: db}
}

// Track user based on given latitude and longitude
func (t *Tracker) Track(ctx context.Context, userID int, request entity.TrackRequest) (entity.Location, *entity.ApplicationError) {
	track := &usecase.Track{}
	return track.Perform(ctx, userID, request, t)
}

// Create locations data
func (t *Tracker) Create(ctx context.Context, locationInsertable entity.LocationInsertable) (int, *entity.ApplicationError) {
	create := &usecase.Create{}
	return create.Perform(ctx, locationInsertable, t.db)
}

// Find locations data based on id
func (t *Tracker) Find(ctx context.Context, locationID int) (entity.Location, *entity.ApplicationError) {
	find := &usecase.Find{}
	return find.Perform(ctx, locationID, t.db)
}

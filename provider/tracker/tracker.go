package tracker

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/tracker/api"
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

// FabricateAPI fabricating tracker related API
func (t *Tracker) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewTrack(t))
}

// Track user based on given latitude and longitude
func (t *Tracker) Track(ctx provider.Context, userID int, request entity.TrackRequest) (entity.Location, *entity.ApplicationError) {
	track := &usecase.Track{}
	return track.Perform(ctx, userID, request, t)
}

// Create locations data
func (t *Tracker) Create(ctx provider.Context, locationInsertable entity.LocationInsertable) (int, *entity.ApplicationError) {
	create := &usecase.Create{}
	return create.Perform(ctx, locationInsertable, t.db)
}

// Find locations data based on id
func (t *Tracker) Find(ctx provider.Context, locationID int) (entity.Location, *entity.ApplicationError) {
	find := &usecase.Find{}
	return find.Perform(ctx, locationID, t.db)
}

package usecase

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Track ...
type Track struct{}

// Perform track logic
func (t *Track) Perform(ctx provider.Context, userID int, request entity.TrackRequest, tracker provider.Tracker) (entity.Location, *entity.ApplicationError) {
	ID, err := tracker.Create(ctx, entity.LocationInsertable{
		Lat:    request.Lat,
		Long:   request.Long,
		UserID: userID,
	})
	if err != nil {
		return entity.Location{}, err
	}

	location, err := tracker.Find(ctx, ID)
	if err != nil {
		return entity.Location{}, err
	}

	return location, nil
}

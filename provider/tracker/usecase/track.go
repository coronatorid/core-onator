package usecase

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("tracker").Str("track")).
			Msg("error when creating tracker")
		return entity.Location{}, err
	}

	location, err := tracker.Find(ctx, ID)
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("tracker").Str("track")).
			Msg("error when finding tracker data")
		return entity.Location{}, err
	}

	return location, nil
}

package usecase

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Create ...
type Create struct{}

// Perform ...
func (t *Create) Perform(ctx provider.Context, locationInsertable entity.LocationInsertable, db provider.DB) (int, *entity.ApplicationError) {
	result, err := db.ExecContext(ctx, "location-create", "insert into locations (`user_id`, `lat`, `long`, `created_at`, `updated_at`) values(?, ?, ?, now(), now())", locationInsertable.UserID, locationInsertable.Lat, locationInsertable.Long)
	if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("tracker").Str("create")).
			Msg("error when creating tracker")

		return 0, util.CreateInternalServerError(ctx)
	}

	ID, err := result.LastInsertId()
	if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("tracker").Str("create")).
			Msg("error when creating tracker")

		return 0, util.CreateInternalServerError(ctx)
	}

	return int(ID), nil
}

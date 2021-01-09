package usecase

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Delete false positive reported cases
type Delete struct{}

// Perform delete report logic
func (d *Delete) Perform(ctx provider.Context, ID int, db provider.DB) *entity.ApplicationError {
	_, err := db.ExecContext(ctx, "delete-reported-cases", "delete from reported_cases where id = ?", ID)
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("delete")).
			Msg("error when deleting reported cases")
		return util.CreateInternalServerError(ctx)
	}

	return nil
}

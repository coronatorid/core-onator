package usecase

import (
	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// UpdateState of reported cases
type UpdateState struct{}

// Perform update state logic
func (u *UpdateState) Perform(ctx provider.Context, state constant.ReportedCasesStatus, ID int, db provider.DB) *entity.ApplicationError {
	_, err := db.ExecContext(ctx, "delete-reported-cases", "update reported_cases set state = ? where id = ?", state, ID)
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("update_state")).
			Msg("error when deleting reported cases")
		return util.CreateInternalServerError(ctx)
	}

	return nil
}

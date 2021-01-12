package usecase

import (
	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Count all reported cases in databases
type Count struct{}

// Perform to count reported cases in databases
func (c *Count) Perform(ctx provider.Context, status constant.ReportedCasesStatus, db provider.DB) (int, *entity.ApplicationError) {
	var total int

	rows := db.QueryRowContext(ctx, "list-report", "select count(id) where status = ?", status)
	if err := rows.Scan(&total); err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Stack().
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("count")).
			Msg("failed query context")
		return total, util.CreateServiceUnavailable(ctx)
	}

	return total, nil
}

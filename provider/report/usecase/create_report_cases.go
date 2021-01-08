package usecase

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// CreateReportCases data
type CreateReportCases struct{}

// Perform logic to new reported cases
func (c *CreateReportCases) Perform(ctx provider.Context, insertable entity.ReportInsertable, tx provider.TX) (int, *entity.ApplicationError) {
	result, err := tx.ExecContext(ctx, "reported-cases-create", "insert into users (user_id, status, image_path, image_deleted, created_at, updated_at) values(?, 2, ?, 0, now(), now())", insertable.UserID, insertable.ImagePath)
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("create_report_cases")).
			Msg("error when creating report cases")
		return 0, util.CreateInternalServerError(ctx)
	}

	ID, err := result.LastInsertId()
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("create_report_cases")).
			Msg("error when getting last inserted id")
		return 0, util.CreateInternalServerError(ctx)
	}

	return int(ID), nil
}

package usecase

import (
	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// List return list of report that available in databases
type List struct{}

// Perform logic to get list of report from databases
func (c *List) Perform(ctx provider.Context, status constant.ReportedCasesStatus, requestMeta entity.RequestMeta, db provider.DB) ([]entity.ReportedCases, *entity.ApplicationError) {
	var reportedCases []entity.ReportedCases

	rows, err := db.QueryContext(ctx, "list-report", "select id, user_id, status, image_path, image_deleted, created_at, updated_at where status = ? from reported_cases limit ?, ?", status, requestMeta.Offset, requestMeta.Limit)
	if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Stack().
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("list")).
			Msg("failed find")
		return reportedCases, util.CreateServiceUnavailable(ctx)
	}

	for rows.Next() {
		var reportedCase entity.ReportedCases

		err := rows.Scan(
			&reportedCase.ID,
			&reportedCase.UserID,
			&reportedCase.Status,
			&reportedCase.ImagePath,
			&reportedCase.ImageDeleted,
			&reportedCase.CreatedAt,
			&reportedCase.UpdatedAt,
		)
		if err != nil {
			log.Error().
				Err(err).
				Str("request_id", util.GetRequestID(ctx)).
				Stack().
				Array("tags", zerolog.Arr().Str("provider").Str("report").Str("list")).
				Msg("failed scan")
			return reportedCases, util.CreateServiceUnavailable(ctx)
		}

		reportedCases = append(reportedCases, reportedCase)
	}

	return reportedCases, nil
}

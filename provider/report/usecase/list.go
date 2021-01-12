package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/codefluence-x/aurelia"
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
	var reportedCases = []entity.ReportedCases{}

	rows, err := db.QueryContext(ctx, "list-report", "select id, user_id, status, image_path, image_deleted, created_at, updated_at from reported_cases where status = ? limit ?, ?", status, requestMeta.Offset, requestMeta.Limit)
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

		expiresAt := time.Now().Add(60 * time.Second).Unix()
		signature := aurelia.Hash(os.Getenv("APP_ENCRIPTION_KEY"), fmt.Sprintf("%d%s", expiresAt, reportedCase.ImagePath))
		reportedCase.ImagePath = fmt.Sprintf("%s%s?signature=%s&expires_at=%d", os.Getenv("API_HOST"), reportedCase.ImagePath, signature, expiresAt)

		reportedCases = append(reportedCases, reportedCase)
	}

	return reportedCases, nil
}

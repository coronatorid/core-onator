package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/codefluence-x/aurelia"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Find find reported cases by user id
type Find struct{}

// Perform find operation
func (f *Find) Perform(ctx provider.Context, ID int, db provider.DB) (entity.ReportedCases, *entity.ApplicationError) {
	var reportedCase entity.ReportedCases

	row := db.QueryRowContext(ctx, "find-reported-cases", "select id, user_id, status, image_path, image_deleted, created_at, updated_at from reported_cases where id = ?", ID)
	if err := row.Scan(
		&reportedCase.ID,
		&reportedCase.UserID,
		&reportedCase.Status,
		&reportedCase.ImagePath,
		&reportedCase.ImageDeleted,
		&reportedCase.CreatedAt,
		&reportedCase.UpdatedAt,
	); err == provider.ErrDBNotFound {
		return reportedCase, &entity.ApplicationError{
			Err:        []error{errors.New("reported cases not found")},
			HTTPStatus: http.StatusNotFound,
		}
	} else if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Stack().
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("find")).
			Msg("failed find")
		return reportedCase, util.CreateServiceUnavailable(ctx)
	}

	expiresAt := time.Now().Add(60 * time.Second).Unix()
	signature := aurelia.Hash(os.Getenv("APP_ENCRIPTION_KEY"), fmt.Sprintf("%d%s", expiresAt, reportedCase.ImagePath))
	reportedCase.ImagePath = fmt.Sprintf("%s%s?signature=%s&expires_at=%d", os.Getenv("API_HOST"), reportedCase.ImagePath, signature, expiresAt)

	return reportedCase, nil
}

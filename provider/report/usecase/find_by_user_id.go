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

// FindByUserID find reported cases by user id
type FindByUserID struct{}

// Perform find operation
func (f *FindByUserID) Perform(ctx provider.Context, userID int, db provider.DB) (entity.ReportedCases, *entity.ApplicationError) {
	var reportedCases entity.ReportedCases

	row := db.QueryRowContext(ctx, "find-reported-cases", "select id, user_id, status, image_path, image_deleted, created_at, updated_at from reported_cases where user_id = ?", userID)
	if err := row.Scan(
		&reportedCases.ID,
		&reportedCases.UserID,
		&reportedCases.Status,
		&reportedCases.ImagePath,
		&reportedCases.ImageDeleted,
		&reportedCases.CreatedAt,
		&reportedCases.UpdatedAt,
	); err == provider.ErrDBNotFound {
		return reportedCases, &entity.ApplicationError{
			Err:        []error{errors.New("reported cases not found")},
			HTTPStatus: http.StatusNotFound,
		}
	} else if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Stack().
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("find_by_user_id")).
			Msg("failed find")
		return reportedCases, &entity.ApplicationError{
			Err:        []error{errors.New("service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	expiresAt := time.Now().Add(60 * time.Second).Unix()
	signature := aurelia.Hash(os.Getenv("APP_ENCRIPTION_KEY"), fmt.Sprintf("%d%s", expiresAt, reportedCases.ImagePath))
	reportedCases.ImagePath = fmt.Sprintf("%s%s?signature=%s&expires_at=%d", os.Getenv("API_HOST"), reportedCases.ImagePath, signature, expiresAt)

	return reportedCases, nil
}

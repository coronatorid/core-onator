package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Find location data by user id
type Find struct{}

// Perform find location by user id
func (f *Find) Perform(ctx context.Context, locationID int, db provider.DB) (entity.Location, *entity.ApplicationError) {
	var location entity.Location

	row := db.QueryRowContext(ctx, "find-location", "select `id`, `user_id`, `lat`, `long`, `created_at`, `updated_at` from locations where id = ?", locationID)
	err := row.Scan(&location.ID, &location.UserID, &location.Lat, &location.Long, &location.CreatedAt, &location.UpdatedAt)
	if err == provider.ErrDBNotFound {
		return location, &entity.ApplicationError{
			Err:        []error{errors.New("location not found")},
			HTTPStatus: http.StatusNotFound,
		}
	} else if err != nil {
		return location, &entity.ApplicationError{
			Err:        []error{errors.New("service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	return location, nil
}

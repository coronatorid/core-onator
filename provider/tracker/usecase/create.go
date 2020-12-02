package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Create ...
type Create struct{}

// Perform ...
func (t *Create) Perform(ctx context.Context, locationInsertable entity.LocationInsertable, db provider.DB) (int, *entity.ApplicationError) {
	result, err := db.ExecContext(ctx, "location-create", "insert into locations (`user_id`, `lat`, `long`, `created_at`, `updated_at`) values(?, ?, ?, now(), now())", locationInsertable.UserID, locationInsertable.Lat, locationInsertable.Long)
	if err != nil {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("internal server error when acquiring last inserted id")},
			HTTPStatus: http.StatusInternalServerError,
		}
	}

	return int(ID), nil
}

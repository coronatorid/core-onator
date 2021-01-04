package usecase

import (
	"errors"
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// CreateReportCases data
type CreateReportCases struct{}

// Perform logic to new reported cases
func (c *CreateReportCases) Perform(ctx provider.Context, insertable entity.ReportInsertable, tx provider.TX) (int, *entity.ApplicationError) {
	result, err := tx.ExecContext(ctx, "reported-cases-create", "insert into users (user_id, status, image_path, image_deleted, created_at, updated_at) values(?, 2, ?, 0, now(), now())", insertable.UserID, insertable.ImagePath)
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

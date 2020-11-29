package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// FindByPhoneNumber user based on given phone number
type FindByPhoneNumber struct{}

// Perform finding user
func (f *FindByPhoneNumber) Perform(ctx context.Context, phoneNumber string, db provider.DB) (entity.User, *entity.ApplicationError) {
	var user entity.User

	row := db.QueryRowContext(ctx, "find-user", "select id, phone, state, created_at, updated_at from users where phone = ?", phoneNumber)
	if err := row.Scan(
		&user.ID,
		&user.Phone,
		&user.State,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err == provider.ErrDBNotFound {
		return user, &entity.ApplicationError{
			Err:        []error{errors.New("user not found")},
			HTTPStatus: http.StatusNotFound,
		}
	} else if err != nil {
		return user, &entity.ApplicationError{
			Err:        []error{errors.New("service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	return user, nil
}

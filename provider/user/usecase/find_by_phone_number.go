package usecase

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// FindByPhoneNumber user based on given phone number
type FindByPhoneNumber struct{}

// Perform finding user
func (f *FindByPhoneNumber) Perform(ctx provider.Context, phoneNumber string, db provider.DB) (entity.User, *entity.ApplicationError) {
	// We should protect our user no matter what
	h := sha256.New()
	_, _ = h.Write([]byte(fmt.Sprintf("%sXXX%s", phoneNumber, os.Getenv("APP_ENCRIPTION_KEY"))))

	var user entity.User

	row := db.QueryRowContext(ctx, "find-user", "select id, phone, state, created_at, updated_at from users where phone = ?", fmt.Sprintf("%x", h.Sum(nil)))
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

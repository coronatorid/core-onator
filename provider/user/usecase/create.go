package usecase

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Create new user
type Create struct{}

// Perform logic to create new user
func (c *Create) Perform(ctx context.Context, userInsertable entity.UserInsertable, db provider.DB) (int, *entity.ApplicationError) {
	// We should protect our user no matter what
	h := sha256.New()
	_, _ = h.Write([]byte(fmt.Sprintf("%sXXX%s", userInsertable.PhoneNumber, os.Getenv("APP_ENCRIPTION_KEY"))))
	phoneNumber := fmt.Sprintf("%x", h.Sum(nil))

	result, err := db.ExecContext(ctx, "user-create", "insert into users (phone, state, created_at, updated_at) values(?, 1, now(), now())", phoneNumber)
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

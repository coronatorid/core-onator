package usecase

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Create new user
type Create struct{}

// Perform logic to create new user
func (c *Create) Perform(ctx provider.Context, userInsertable entity.UserInsertable, db provider.DB) (int, *entity.ApplicationError) {
	// We should protect our user no matter what
	h := sha256.New()
	_, _ = h.Write([]byte(fmt.Sprintf("%sXXX%s", userInsertable.PhoneNumber, os.Getenv("APP_ENCRIPTION_KEY"))))
	phoneNumber := fmt.Sprintf("%x", h.Sum(nil))

	result, err := db.ExecContext(ctx, "user-create", "insert into users (phone, state, created_at, updated_at) values(?, 1, now(), now())", phoneNumber)
	if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("user").Str("create")).
			Msg("error when creating user")
		return 0, util.CreateInternalServerError(ctx)
	}

	ID, err := result.LastInsertId()
	if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("user").Str("create")).
			Msg("error when creating user")
		return 0, util.CreateInternalServerError(ctx)
	}

	return int(ID), nil
}

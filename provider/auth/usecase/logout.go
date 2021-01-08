package usecase

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logout logic from altair application
type Logout struct {
}

// Perform logout
func (l *Logout) Perform(ctx provider.Context, request entity.RevokeTokenRequest, altair provider.Altair) *entity.ApplicationError {
	if err := altair.RevokeToken(ctx, request); err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("auth").Str("logout")).
			Msg("error when revoking altair token")
		return err
	}

	return nil
}

package usecase

import (
	"os"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// DeleteFile ...
type DeleteFile struct{}

// Perform delete file
func (d *DeleteFile) Perform(ctx provider.Context, filePath string) *entity.ApplicationError {
	err := os.Remove(filePath)
	if err != nil {
		log.Error().
			Err(err).
			Stack().
			Str("request_id", util.GetRequestID(ctx)).
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("delete_file")).
			Msg("error when deleting reported cases file")
		return util.CreateInternalServerError(ctx)
	}

	return nil
}

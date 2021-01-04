package usecase

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// UploadFileAndCreate data for reported cases
type UploadFileAndCreate struct{}

// Perform logic to new reported cases
func (c *UploadFileAndCreate) Perform(ctx provider.Context, db provider.DB) (int, *entity.ApplicationError) {

	return 0, nil
}

package usecase

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Logout logic from altair application
type Logout struct {
}

// Perform logout
func (l *Logout) Perform(ctx provider.Context, request entity.RevokeTokenRequest, altair provider.Altair) *entity.ApplicationError {
	return altair.RevokeToken(ctx, request)
}

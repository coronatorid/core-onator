package provider

import (
	"context"

	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./altair.go -destination=./mocks/altair_mock.go -package mockProvider

// Altair interface connector
type Altair interface {
	GrantToken(ctx context.Context, request entity.GrantTokenRequest) (entity.OauthAccessToken, *entity.ApplicationError)
}
